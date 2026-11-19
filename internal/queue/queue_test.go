package queue

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestQueueIsFIFO(t *testing.T) {
	q := New(4)
	for _, id := range []string{"a", "b", "c"} {
		if err := q.Push(Event{ID: id}); err != nil {
			t.Fatalf("push %s: %v", id, err)
		}
	}

	for _, id := range []string{"a", "b", "c"} {
		got, err := q.Pop(context.Background())
		if err != nil {
			t.Fatalf("pop: %v", err)
		}
		if got.ID != id {
			t.Fatalf("popped %q, want %q", got.ID, id)
		}
	}
}

func TestPushRejectsWhenFull(t *testing.T) {
	q := New(1)
	if err := q.Push(Event{ID: "a"}); err != nil {
		t.Fatalf("push: %v", err)
	}
	if err := q.Push(Event{ID: "b"}); !errors.Is(err, ErrFull) {
		t.Fatalf("error = %v, want ErrFull", err)
	}
}

func TestPopHonoursContext(t *testing.T) {
	q := New(2)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	if _, err := q.Pop(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v", err)
	}
}

func TestPushStampsEnqueuedAt(t *testing.T) {
	q := New(2)
	if err := q.Push(Event{ID: "a"}); err != nil {
		t.Fatalf("push: %v", err)
	}

	got, _ := q.Pop(context.Background())
	if got.EnqueuedAt.IsZero() {
		t.Fatal("EnqueuedAt was not stamped")
	}
}

func TestClosePreservesBufferedEvents(t *testing.T) {
	q := New(1)
	if err := q.Push(Event{ID: "buffered"}); err != nil {
		t.Fatal(err)
	}
	q.Close()
	q.Close()
	if err := q.Push(Event{ID: "late"}); !errors.Is(err, ErrClosed) {
		t.Fatalf("push error = %v, want ErrClosed", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	event, err := q.Pop(ctx)
	if err != nil || event.ID != "buffered" {
		t.Fatalf("pop = %v, %v, want buffered event", event, err)
	}
	if _, err := q.Pop(ctx); !errors.Is(err, ErrClosed) {
		t.Fatalf("pop error = %v, want ErrClosed", err)
	}
}

func TestConcurrentPushAndClose(t *testing.T) {
	for range 1000 {
		q := New(1)
		start := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			<-start
			if err := q.Push(Event{ID: "event"}); err != nil && !errors.Is(err, ErrClosed) {
				t.Errorf("push error = %v, want nil or ErrClosed", err)
			}
		}()
		for range 2 {
			go func() {
				defer wg.Done()
				<-start
				q.Close()
			}()
		}
		close(start)
		wg.Wait()
	}
}
