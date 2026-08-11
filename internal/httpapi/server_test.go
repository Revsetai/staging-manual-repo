package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/breaker"
	"github.com/Revsetai/staging-manual-repo/internal/config"
	"github.com/Revsetai/staging-manual-repo/internal/metrics"
)

func newTestServer() (*Server, http.Handler) {
	s := New(config.Default(), metrics.NewRegistry())
	return s, s.Handler()
}

func get(t *testing.T, h http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestHealthzIsAlwaysOK(t *testing.T) {
	_, h := newTestServer()
	rec := get(t, h, "/healthz")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("body is not JSON: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("status field = %v, want ok", body["status"])
	}
}

func TestReadyzFailsWhenACheckFails(t *testing.T) {
	s, h := newTestServer()
	s.AddReadinessCheck("upstream", func(context.Context) error {
		return errors.New("dial tcp: refused")
	})

	rec := get(t, h, "/readyz")
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "dial tcp: refused") {
		t.Errorf("body should surface the check error: %s", rec.Body.String())
	}
}

func TestReadyzFailsWhileDraining(t *testing.T) {
	s, h := newTestServer()
	s.Drain()

	rec := get(t, h, "/readyz")
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", rec.Code)
	}
}

func TestRequestsAreCountedPerRoute(t *testing.T) {
	s, h := newTestServer()
	s.AddReadinessCheck("upstream", func(context.Context) error {
		return errors.New("still connecting")
	})

	get(t, h, "/healthz")
	get(t, h, "/healthz")
	get(t, h, "/readyz")

	rec := get(t, h, "/metrics")
	if got := rec.Header().Get("Content-Type"); got != metrics.ContentType {
		t.Errorf("content type = %q", got)
	}

	body := rec.Body.String()
	for _, want := range []string{
		`ingestd_http_requests_total{route="/healthz",status="2xx"} 2`,
		`ingestd_http_requests_total{route="/readyz",status="5xx"} 1`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("output is missing %q\n---\n%s", want, body)
		}
	}
	// The scrape is counted after its own body is written, so it has no
	// request series of its own on this first pass.
	if strings.Contains(body, `ingestd_http_requests_total{route="/metrics"`) {
		t.Errorf("the in-flight scrape should not count itself:\n%s", body)
	}
}

func TestResponseBytesAreRecorded(t *testing.T) {
	_, h := newTestServer()
	first := get(t, h, "/healthz")

	body := get(t, h, "/metrics").Body.String()
	want := "ingestd_http_response_bytes_total{route=\"/healthz\"} " +
		strconv.Itoa(first.Body.Len())
	if !strings.Contains(body, want) {
		t.Errorf("output is missing %q\n---\n%s", want, body)
	}
}

func TestBuildInfoIsExported(t *testing.T) {
	_, h := newTestServer()

	body := get(t, h, "/metrics").Body.String()
	want := `ingestd_build_info{env="dev",version="dev"} 1`
	if !strings.Contains(body, want) {
		t.Errorf("output is missing %q\n---\n%s", want, body)
	}
	if !strings.Contains(body, "# TYPE ingestd_build_info gauge") {
		t.Errorf("build_info should be exposed as a gauge:\n%s", body)
	}
}

func TestWatchBreakerFailsReadinessWhenOpen(t *testing.T) {
	s, h := newTestServer()
	b := breaker.New(breaker.Settings{FailureThreshold: 1, Cooldown: time.Hour})
	s.WatchBreaker("upstream", b)

	if rec := get(t, h, "/readyz"); rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 while the circuit is closed", rec.Code)
	}

	_ = b.Call(context.Background(), func(context.Context) error {
		return errors.New("upstream down")
	})

	rec := get(t, h, "/readyz")
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503 once the circuit opens", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "circuit open") {
		t.Errorf("body should explain the failure: %s", rec.Body.String())
	}

	body := get(t, h, "/metrics").Body.String()
	if !strings.Contains(body, `ingestd_breaker_open{dependency="upstream"} 1`) {
		t.Errorf("the open circuit should show on the gauge:\n%s", body)
	}
}

func TestStatusClassBuckets(t *testing.T) {
	cases := map[int]string{200: "2xx", 204: "2xx", 301: "3xx", 404: "4xx", 503: "5xx"}
	for status, want := range cases {
		if got := statusClass(status); got != want {
			t.Errorf("statusClass(%d) = %q, want %q", status, got, want)
		}
	}
}
