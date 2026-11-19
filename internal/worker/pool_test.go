package worker

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
)

func quietOptions(size int) Options {
	return Options{
		Size:           size,
		HandlerTimeout: time.Second,
		ShutdownGrace:  time.Second,
		Logger:         slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

func TestPoolProcessesEveryEvent(t *testing.T) {
	q := queue.New(32)
	var mu sync.Mutex
	seen := map[string]bool{}

	p := NewPool(q, func(ctx context.Context, e queue.Event) error {
		mu.Lock()
		seen[e.ID] = true
		mu.Unlock()
		return nil
	}, quietOptions(3))

	if err := p.Start(context.Background()); err != nil {
		t.Fatalf("start: %v", err)
	}
	for _, id := range []string{"a", "b", "c", "d"} {
		if err := q.Push(queue.Event{ID: id, Priority: queue.PriorityNormal}); err != nil {
			t.Fatalf("push %s: %v", id, err)
		}
	}

	waitFor(t, func() bool { return p.Stats().Processed == 4 })
	if err := p.Stop(); err != nil {
		t.Fatalf("stop: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(seen) != 4 {
		t.Fatalf("saw %d distinct events, want 4", len(seen))
	}
}

func TestPoolCountsFailures(t *testing.T) {
	q := queue.New(8)
	p := NewPool(q, func(ctx context.Context, e queue.Event) error {
		if e.ID == "bad" {
			return errors.New("handler exploded")
		}
		return nil
	}, quietOptions(1))

	_ = p.Start(context.Background())
	_ = q.Push(queue.Event{ID: "good"})
	_ = q.Push(queue.Event{ID: "bad"})

	waitFor(t, func() bool { return p.Stats().Processed == 2 })
	_ = p.Stop()

	if got := p.Stats().Failed; got != 1 {
		t.Fatalf("failed = %d, want 1", got)
	}
}

func TestStartTwiceIsRejected(t *testing.T) {
	q := queue.New(4)
	p := NewPool(q, func(context.Context, queue.Event) error { return nil }, quietOptions(1))

	_ = p.Start(context.Background())
	defer p.Stop()

	if err := p.Start(context.Background()); !errors.Is(err, ErrAlreadyRunning) {
		t.Fatalf("error = %v, want ErrAlreadyRunning", err)
	}
}

func TestPoolTracksLatency(t *testing.T) {
	q := queue.New(8)
	p := NewPool(q, func(ctx context.Context, e queue.Event) error {
		if e.ID == "slow" {
			time.Sleep(25 * time.Millisecond)
		}
		return nil
	}, quietOptions(1))

	_ = p.Start(context.Background())
	_ = q.Push(queue.Event{ID: "quick"})
	_ = q.Push(queue.Event{ID: "slow"})

	waitFor(t, func() bool { return p.Stats().Processed == 2 })
	_ = p.Stop()

	stats := p.Stats()
	if stats.SlowestLatency < 25*time.Millisecond {
		t.Errorf("slowest latency = %s, want at least 25ms", stats.SlowestLatency)
	}
	if stats.MeanLatency() <= 0 || stats.MeanLatency() > stats.SlowestLatency {
		t.Errorf("mean latency = %s, slowest = %s", stats.MeanLatency(), stats.SlowestLatency)
	}
}

func TestStateTracksLifecycle(t *testing.T) {
	q := queue.New(4)
	p := NewPool(q, func(context.Context, queue.Event) error { return nil }, quietOptions(1))

	if got := p.State(); got != "idle" {
		t.Fatalf("state = %q, want idle", got)
	}
	_ = p.Start(context.Background())
	if got := p.State(); got != "running" {
		t.Fatalf("state = %q, want running", got)
	}
	_ = p.Stop()
	if got := p.State(); got != "stopped" {
		t.Fatalf("state = %q, want stopped", got)
	}
}

func TestMeanLatencyOfAnIdlePoolIsZero(t *testing.T) {
	if got := (Stats{}).MeanLatency(); got != 0 {
		t.Fatalf("mean latency = %s, want 0", got)
	}
}

func TestStopReportsAbandonedEvents(t *testing.T) {
	q := queue.New(16)

	// A handler that only returns when its context is cancelled: the pool
	// has to hit the shutdown grace and tear the worker down mid-event.
	opts := quietOptions(1)
	opts.ShutdownGrace = 50 * time.Millisecond
	p := NewPool(q, func(ctx context.Context, e queue.Event) error {
		<-ctx.Done()
		return ctx.Err()
	}, opts)

	_ = p.Start(context.Background())
	for i := 0; i < 5; i++ {
		_ = q.Push(queue.Event{ID: "e" + string(rune('0'+i))})
	}
	waitFor(t, func() bool { return p.Stats().InFlight == 1 })

	err := p.Stop()
	if err == nil {
		t.Fatal("Stop should report that it ran out of grace")
	}
	if !strings.Contains(err.Error(), "still in flight") {
		t.Errorf("error = %v, want it to name the in-flight work", err)
	}
	if got := p.Stats().Abandoned; got != 4 {
		t.Fatalf("abandoned = %d, want 4", got)
	}
}

func TestGracefulStopDrainsEverything(t *testing.T) {
	q := queue.New(16)
	p := NewPool(q, func(context.Context, queue.Event) error { return nil }, quietOptions(2))

	_ = p.Start(context.Background())
	for i := 0; i < 5; i++ {
		_ = q.Push(queue.Event{ID: "e" + string(rune('0'+i))})
	}
	waitFor(t, func() bool { return p.Stats().Processed == 5 })

	if err := p.Stop(); err != nil {
		t.Fatalf("stop: %v", err)
	}
	if got := p.Stats().Abandoned; got != 0 {
		t.Fatalf("abandoned = %d, want 0", got)
	}
}

func TestStopIsIdempotent(t *testing.T) {
	q := queue.New(4)
	p := NewPool(q, func(context.Context, queue.Event) error { return nil }, quietOptions(1))
	_ = p.Start(context.Background())

	if err := p.Stop(); err != nil {
		t.Fatalf("first stop: %v", err)
	}
	if err := p.Stop(); err != nil {
		t.Fatalf("second stop: %v", err)
	}
}

func waitFor(t *testing.T, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("condition was not met before the deadline")
}
