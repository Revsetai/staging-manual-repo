// Package httpapi serves the daemon's operational API and dashboard.
package httpapi

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/operations"
	"github.com/Revsetai/staging-manual-repo/internal/queue"
)

const maxEventBody = 64 << 10

// EventSink is implemented by queues that accept events.
type EventSink interface {
	Push(queue.Event) error
}

// SnapshotProvider supplies the API with a complete runtime snapshot.
type SnapshotProvider interface {
	Snapshot(context.Context) operations.Snapshot
}

// Server holds the handlers for the operations surface.
type Server struct {
	events    EventSink
	snapshots SnapshotProvider
	started   time.Time
	nextID    atomic.Uint64
}

//go:embed web/index.html web/styles.css web/app.mjs
var dashboardFiles embed.FS

// New builds a Server over the live runtime dependencies.
func New(events EventSink, snapshots SnapshotProvider) *Server {
	return &Server{events: events, snapshots: snapshots, started: time.Now()}
}

// Handler returns the API and dashboard mux.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /api/snapshot", s.handleSnapshot)
	mux.HandleFunc("POST /api/events", s.handleEvent)
	mux.Handle("GET /", s.dashboardHandler())
	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "ok uptime=%s\n", time.Since(s.started).Round(time.Second))
}

func (s *Server) handleSnapshot(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.snapshots.Snapshot(r.Context()))
}

type eventRequest struct {
	ID      string `json:"id"`
	Source  string `json:"source"`
	Payload string `json:"payload"`
}

func (s *Server) handleEvent(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxEventBody)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request eventRequest
	if err := decoder.Decode(&request); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid_event", err.Error())
		return
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeAPIError(w, http.StatusBadRequest, "invalid_event", "request body must contain one JSON object")
		return
	}

	request.ID = strings.TrimSpace(request.ID)
	request.Source = strings.TrimSpace(request.Source)
	if request.Source == "" {
		writeAPIError(w, http.StatusBadRequest, "invalid_event", "source is required")
		return
	}
	if request.ID == "" {
		request.ID = fmt.Sprintf("dashboard-%06d", s.nextID.Add(1))
	}

	event := queue.Event{
		ID:      request.ID,
		Source:  request.Source,
		Payload: []byte(request.Payload),
	}
	if err := s.events.Push(event); err != nil {
		switch {
		case errors.Is(err, queue.ErrFull):
			writeAPIError(w, http.StatusTooManyRequests, "queue_full", "the queue is at capacity")
		case errors.Is(err, queue.ErrClosed):
			writeAPIError(w, http.StatusServiceUnavailable, "queue_closed", "the pipeline is stopped")
		default:
			writeAPIError(w, http.StatusInternalServerError, "enqueue_failed", "the event could not be queued")
		}
		return
	}

	writeJSON(w, http.StatusAccepted, map[string]string{"id": event.ID, "status": "queued"})
}

func (s *Server) dashboardHandler() http.Handler {
	web, err := fs.Sub(dashboardFiles, "web")
	if err != nil {
		panic(fmt.Sprintf("dashboard assets: %v", err))
	}
	return http.FileServer(http.FS(web))
}

// Uptime reports how long the server has been up.
func (s *Server) Uptime() time.Duration { return time.Since(s.started) }

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeAPIError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]string{"code": code, "message": message})
}
