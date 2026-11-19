// Package httpapi exposes ingestd's operational HTTP surface: liveness,
// readiness, and the metrics scrape endpoint.
package httpapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/breaker"
	"github.com/Revsetai/staging-manual-repo/internal/config"
	"github.com/Revsetai/staging-manual-repo/internal/metrics"
)

// ReadinessCheck reports whether one subsystem is ready to serve.
type ReadinessCheck func(ctx context.Context) error

// Server wires the admin handlers onto a mux.
type Server struct {
	cfg      config.Config
	registry *metrics.Registry
	encoder  *metrics.Encoder
	checks   map[string]ReadinessCheck
	started  time.Time
	draining atomic.Bool
}

// New builds a Server for the given configuration and registry.
func New(cfg config.Config, registry *metrics.Registry) *Server {
	encoder := metrics.NewEncoder("ingestd")
	encoder.Describe("http_requests_total", "Admin HTTP requests by route and status class.")
	encoder.Describe("http_response_bytes_total", "Bytes written by route.")
	encoder.Describe("uptime_seconds", "Seconds since the process started serving.")
	encoder.Describe("build_info", "Build and deployment identity; the value is always 1.")
	encoder.Describe("breaker_open", "1 when a dependency's circuit is open, 0 otherwise.")

	// Registered once at construction: an info series is constant, so it
	// belongs beside the other declarations rather than in the scrape path.
	registry.Info("build_info", metrics.Labels{
		"version": Version,
		"env":     cfg.Environment.String(),
	})

	return &Server{
		cfg:      cfg,
		registry: registry,
		encoder:  encoder,
		checks:   make(map[string]ReadinessCheck),
		started:  time.Now(),
	}
}

// AddReadinessCheck registers a named readiness probe.
func (s *Server) AddReadinessCheck(name string, check ReadinessCheck) {
	s.checks[name] = check
}

// WatchBreaker reports the named breaker's health as a readiness check and
// mirrors its state onto a gauge, so an open circuit is visible both to the
// orchestrator and on the dashboard.
func (s *Server) WatchBreaker(name string, b *breaker.Breaker) {
	gauge := s.registry.Gauge("breaker_open", metrics.Labels{"dependency": name})

	s.AddReadinessCheck(name, func(context.Context) error {
		snap := b.Snapshot()
		if snap.Healthy() {
			gauge.Set(0)
			return nil
		}
		gauge.Set(1)
		return fmt.Errorf("circuit open for %s after %d failures",
			snap.OpenFor.Round(time.Second), snap.Failures)
	})
}

// Drain flips the server into draining mode so readiness starts failing
// before the listener is closed.
func (s *Server) Drain() { s.draining.Store(true) }

// Handler returns the mux serving every admin route.
//
// Each route is wrapped individually so that its metrics carry the route as a
// label. A single wrapper around the mux could only ever report one aggregate
// number, which said nothing about which endpoint was failing.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /healthz", s.instrument("/healthz", s.handleHealth))
	mux.Handle("GET /readyz", s.instrument("/readyz", s.handleReady))
	mux.Handle("GET /metrics", s.instrument("/metrics", s.handleMetrics))
	return mux
}

// instrument counts requests, status classes, and bytes for one route.
func (s *Server) instrument(route string, h http.HandlerFunc) http.Handler {
	bytes := s.registry.Counter("http_response_bytes_total", metrics.Labels{"route": route})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &responseRecorder{ResponseWriter: w, status: http.StatusOK}
		h.ServeHTTP(rec, r)

		s.registry.Counter("http_requests_total", metrics.Labels{
			"route":  route,
			"status": statusClass(rec.status),
		}).Inc()
		bytes.Add(uint64(rec.written))
	})
}

// statusClass buckets a status code the way an alert would group it, keeping
// the label cardinality at four instead of one series per code.
func statusClass(status int) string {
	switch {
	case status >= 500:
		return "5xx"
	case status >= 400:
		return "4xx"
	case status >= 300:
		return "3xx"
	default:
		return "2xx"
	}
}

// responseRecorder captures the status and size of a response.
type responseRecorder struct {
	http.ResponseWriter
	status  int
	written int
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.written += n
	return n, err
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"env":     s.cfg.Environment,
		"uptime":  time.Since(s.started).Round(time.Second).String(),
		"version": Version,
	})
}

func (s *Server) handleReady(w http.ResponseWriter, r *http.Request) {
	if s.draining.Load() {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{
			"status": "draining",
		})
		return
	}

	results := make(map[string]string, len(s.checks))
	status := http.StatusOK
	for name, check := range s.checks {
		if err := check(r.Context()); err != nil {
			results[name] = err.Error()
			status = http.StatusServiceUnavailable
			continue
		}
		results[name] = "ok"
	}

	writeJSON(w, status, map[string]any{
		"status": statusWord(status),
		"checks": results,
	})
}

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	s.registry.Gauge("uptime_seconds", nil).Set(time.Since(s.started).Seconds())

	w.Header().Set("Content-Type", metrics.ContentType)
	if err := s.encoder.EncodeRegistry(w, s.registry); err != nil {
		http.Error(w, "failed to encode metrics", http.StatusInternalServerError)
	}
}

func statusWord(status int) string {
	if status == http.StatusOK {
		return "ready"
	}
	return "unready"
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// Version is stamped at build time via -ldflags.
var Version = "dev"
