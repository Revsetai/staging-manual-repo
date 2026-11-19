package queue

import (
	"context"
	"errors"
	"testing"
	"time"
)

func event(id string, p Priority) Event {
	return Event{ID: id, Source: "test", Priority: p, Payload: []byte(id)}
}

func TestPopReturnsHighestPriorityFirst(t *testing.T) {
	q := New(10)
	mustPush(t, q, event("low", PriorityLow))
	mustPush(t, q, event("high", PriorityHigh))
	mustPush(t, q, event("normal", PriorityNormal))

	want := []string{"high", "normal", "low"}
	for _, id := range want {
		got, err := q.Pop(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.ID != id {
			t.Fatalf("popped %q, want %q", got.ID, id)
		}
	}
}

func TestEqualPriorityDrainsFIFO(t *testing.T) {
	q := New(10)
	for _, id := range []string{"a", "b", "c"} {
		mustPush(t, q, event(id, PriorityNormal))
	}

	for _, id := range []string{"a", "b", "c"} {
		got, _ := q.TryPop()
		if got.ID != id {
			t.Fatalf("popped %q, want %q", got.ID, id)
		}
	}
}

func TestPushRejectsWhenFullAtEqualPriority(t *testing.T) {
	q := New(2)
	mustPush(t, q, event("a", PriorityNormal))
	mustPush(t, q, event("b", PriorityNormal))

	if err := q.Push(event("c", PriorityNormal)); !errors.Is(err, ErrFull) {
		t.Fatalf("error = %v, want ErrFull", err)
	}
	if got := q.Evicted(); got != 0 {
		t.Errorf("evicted = %d, want 0; an equal-priority arrival must not evict", got)
	}
}

func TestHighPriorityEvictsTheLowest(t *testing.T) {
	q := New(3)
	mustPush(t, q, event("low-old", PriorityLow))
	mustPush(t, q, event("low-new", PriorityLow))
	mustPush(t, q, event("normal", PriorityNormal))

	mustPush(t, q, event("control", PriorityHigh))

	if got := q.Len(); got != 3 {
		t.Fatalf("len = %d, want the capacity of 3", got)
	}
	if got := q.Evicted(); got != 1 {
		t.Fatalf("evicted = %d, want 1", got)
	}

	// The newest of the two low-priority events is the one discarded: the
	// older one has already waited, so dropping it wastes more work.
	want := []string{"control", "normal", "low-old"}
	for _, id := range want {
		got, ok := q.TryPop()
		if !ok {
			t.Fatalf("queue drained early, expected %q", id)
		}
		if got.ID != id {
			t.Fatalf("popped %q, want %q", got.ID, id)
		}
	}
}

func TestEvictionKeepsTheHeapOrdered(t *testing.T) {
	q := New(4)
	for i, p := range []Priority{PriorityLow, PriorityLow, PriorityNormal, PriorityLow} {
		mustPush(t, q, event(string(rune('a'+i)), p))
	}

	// Three high-priority arrivals evict all three low-priority events.
	for i := 0; i < 3; i++ {
		mustPush(t, q, event("high"+string(rune('0'+i)), PriorityHigh))
	}

	var order []Priority
	for {
		e, ok := q.TryPop()
		if !ok {
			break
		}
		order = append(order, e.Priority)
	}

	for i := 1; i < len(order); i++ {
		if order[i] > order[i-1] {
			t.Fatalf("drain order is not priority-descending: %v", order)
		}
	}
	if got := q.Evicted(); got != 3 {
		t.Errorf("evicted = %d, want 3", got)
	}
}

func TestPopUnblocksOnClose(t *testing.T) {
	q := New(4)
	done := make(chan error, 1)
	go func() {
		_, err := q.Pop(context.Background())
		done <- err
	}()

	time.Sleep(10 * time.Millisecond)
	q.Close()

	select {
	case err := <-done:
		if !errors.Is(err, ErrClosed) {
			t.Fatalf("error = %v, want ErrClosed", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Pop did not return after Close")
	}
}

func TestCloseIsIdempotent(t *testing.T) {
	q := New(4)
	q.Close()
	q.Close()

	if !q.Closed() {
		t.Fatal("queue should report itself closed")
	}
	if err := q.Push(event("a", PriorityNormal)); !errors.Is(err, ErrClosed) {
		t.Fatalf("error = %v, want ErrClosed", err)
	}
}

func TestBufferedEventsSurviveClose(t *testing.T) {
	q := New(4)
	mustPush(t, q, event("a", PriorityNormal))
	mustPush(t, q, event("b", PriorityHigh))
	q.Close()

	got, err := q.Pop(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "b" {
		t.Fatalf("popped %q, want the buffered high-priority event", got.ID)
	}
	if _, err := q.Pop(context.Background()); err != nil {
		t.Fatalf("second buffered event should still drain: %v", err)
	}
	if _, err := q.Pop(context.Background()); !errors.Is(err, ErrClosed) {
		t.Fatalf("error = %v, want ErrClosed once drained", err)
	}
}

func TestPushAfterHeavyChurnStillSignals(t *testing.T) {
	q := New(2)
	// Drain through TryPop so the wake-up buffer fills without a consumer
	// ever taking a token; a later Push must still wake a blocked Pop.
	for i := 0; i < 8; i++ {
		mustPush(t, q, event("churn", PriorityNormal))
		q.TryPop()
	}

	done := make(chan string, 1)
	go func() {
		e, err := q.Pop(context.Background())
		if err != nil {
			done <- "error: " + err.Error()
			return
		}
		done <- e.ID
	}()

	time.Sleep(10 * time.Millisecond)
	mustPush(t, q, event("late", PriorityNormal))

	select {
	case got := <-done:
		if got != "late" {
			t.Fatalf("Pop returned %q, want late", got)
		}
	case <-time.After(time.Second):
		t.Fatal("Pop was never woken by the late push")
	}
}

func TestPopHonoursContext(t *testing.T) {
	q := New(4)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	if _, err := q.Pop(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want DeadlineExceeded", err)
	}
}

func TestPushStampsEnqueuedAt(t *testing.T) {
	q := New(4)
	mustPush(t, q, event("a", PriorityNormal))

	got, _ := q.TryPop()
	if got.EnqueuedAt.IsZero() {
		t.Fatal("EnqueuedAt was not stamped")
	}
	if age := got.Age(time.Now()); age < 0 {
		t.Fatalf("age = %s, want a non-negative duration", age)
	}
}

func mustPush(t *testing.T, q *Queue, e Event) {
	t.Helper()
	if err := q.Push(e); err != nil {
		t.Fatalf("push %q: %v", e.ID, err)
	}
}
