package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/operations"
	"github.com/Revsetai/staging-manual-repo/internal/queue"
	"github.com/Revsetai/staging-manual-repo/internal/worker"
)

type eventSinkStub struct {
	events []queue.Event
	err    error
}

func (s *eventSinkStub) Push(event queue.Event) error {
	s.events = append(s.events, event)
	return s.err
}

type snapshotProviderStub struct{ snapshot operations.Snapshot }

func (s snapshotProviderStub) Snapshot(context.Context) operations.Snapshot { return s.snapshot }

func newTestServer(sink *eventSinkStub) *Server {
	return New(sink, snapshotProviderStub{snapshot: operations.Snapshot{
		CapturedAt: time.Unix(10, 0),
		UptimeMS:   2500,
		Queue:      queue.Stats{Depth: 2, Capacity: 8},
		Workers:    worker.Stats{Processed: 5, Active: 1},
		Health:     operations.Health{Status: "healthy", Pressure: "low"},
	}})
}

func TestHealthzReturnsOK(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestServer(&eventSinkStub{}).Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.HasPrefix(rec.Body.String(), "ok ") {
		t.Errorf("body = %q", rec.Body.String())
	}
}

func TestSnapshotReturnsRuntimeState(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestServer(&eventSinkStub{}).Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/snapshot", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var got operations.Snapshot
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Queue.Depth != 2 || got.Workers.Processed != 5 {
		t.Fatalf("snapshot = %#v", got)
	}
}

func TestEventIsValidatedAndQueued(t *testing.T) {
	sink := &eventSinkStub{}
	rec := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/events", strings.NewReader(`{"source":"billing","payload":"{\"total\":42}"}`))
	request.Header.Set("Content-Type", "application/json")
	newTestServer(sink).Handler().ServeHTTP(rec, request)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want 202: %s", rec.Code, rec.Body.String())
	}
	if len(sink.events) != 1 || sink.events[0].Source != "billing" {
		t.Fatalf("events = %#v", sink.events)
	}
	if !strings.HasPrefix(sink.events[0].ID, "dashboard-") {
		t.Fatalf("event id = %q", sink.events[0].ID)
	}
}

func TestFullQueueReturnsBackpressureResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	sink := &eventSinkStub{err: queue.ErrFull}
	request := httptest.NewRequest(http.MethodPost, "/api/events", strings.NewReader(`{"source":"demo","payload":"{}"}`))
	newTestServer(sink).Handler().ServeHTTP(rec, request)

	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "queue_full") {
		t.Fatalf("body = %s", rec.Body.String())
	}
}

func TestUnexpectedQueueErrorReturnsServerError(t *testing.T) {
	rec := httptest.NewRecorder()
	sink := &eventSinkStub{err: errors.New("disk unavailable")}
	request := httptest.NewRequest(http.MethodPost, "/api/events", strings.NewReader(`{"source":"demo","payload":"{}"}`))
	newTestServer(sink).Handler().ServeHTTP(rec, request)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rec.Code)
	}
}

func TestDashboardIsEmbedded(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestServer(&eventSinkStub{}).Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Pipeline control room") {
		t.Fatal("dashboard missing from response")
	}
}

func TestUnknownRouteIs404(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestServer(&eventSinkStub{}).Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/readyz", nil))

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}
