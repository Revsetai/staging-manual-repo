package worker

import (
	"context"
	"errors"
	"testing"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
)

func TestPoolHandlesEveryEvent(t *testing.T) {
	events := queue.New(4)
	p := NewPool(events, func(context.Context, queue.Event) error { return nil }, 2)
	p.Start(context.Background())

	for _, id := range []string{"a", "b", "c", "d"} {
		if err := events.Push(queue.Event{ID: id}); err != nil {
			t.Fatalf("push %s: %v", id, err)
		}
	}
	events.Close()
	p.Stop()

	if got := p.Stats(); got.Processed != 4 || got.Active != 0 {
		t.Fatalf("stats = %#v, want 4 processed and none active", got)
	}
}

func TestPoolCountsFailures(t *testing.T) {
	events := queue.New(2)
	p := NewPool(events, func(_ context.Context, event queue.Event) error {
		if event.ID == "bad" {
			return errors.New("boom")
		}
		return nil
	}, 1)
	p.Start(context.Background())

	_ = events.Push(queue.Event{ID: "good"})
	_ = events.Push(queue.Event{ID: "bad"})
	events.Close()
	p.Stop()

	if got := p.Failed(); got != 1 {
		t.Fatalf("failed = %d, want 1", got)
	}
}
