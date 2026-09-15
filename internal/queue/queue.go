// Package queue buffers events between intake and the workers.
package queue

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Event is one unit of work moving through the pipeline.
type Event struct {
	ID         string
	Source     string
	Payload    []byte
	EnqueuedAt time.Time
}

// Stats is a point-in-time view of the queue for operational consumers.
type Stats struct {
	Depth    int  `json:"depth"`
	Capacity int  `json:"capacity"`
	Closed   bool `json:"closed"`
}

var (
	// ErrFull is returned when the buffer is full.
	ErrFull = errors.New("queue: full")
	// ErrClosed is returned once the queue is closed.
	ErrClosed = errors.New("queue: closed")
)

// Queue is a bounded FIFO buffer.
//
// TODO: strict FIFO means a control-plane event queues up behind whatever
// backlog happens to be in front of it, and there is no way to say that one
// event matters more than another.
type Queue struct {
	events chan Event

	mu     sync.Mutex
	closed bool
}

// New builds a queue holding at most capacity events.
func New(capacity int) *Queue {
	if capacity < 1 {
		capacity = 1
	}
	return &Queue{events: make(chan Event, capacity)}
}

// Push enqueues an event, or returns ErrFull if there is no room.
func (q *Queue) Push(e Event) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return ErrClosed
	}

	if e.EnqueuedAt.IsZero() {
		e.EnqueuedAt = time.Now()
	}

	select {
	case q.events <- e:
		return nil
	default:
		return ErrFull
	}
}

// Pop blocks until an event arrives, the queue closes, or ctx is done.
func (q *Queue) Pop(ctx context.Context) (Event, error) {
	select {
	case e, ok := <-q.events:
		if !ok {
			return Event{}, ErrClosed
		}
		return e, nil
	case <-ctx.Done():
		return Event{}, ctx.Err()
	}
}

// Len reports how many events are buffered.
func (q *Queue) Len() int { return len(q.events) }

// Cap reports the queue's capacity.
func (q *Queue) Cap() int { return cap(q.events) }

// Snapshot reports queue pressure without exposing the backing channel.
func (q *Queue) Snapshot() Stats {
	q.mu.Lock()
	defer q.mu.Unlock()

	return Stats{
		Depth:    len(q.events),
		Capacity: cap(q.events),
		Closed:   q.closed,
	}
}

// Close shuts the queue and is safe to call more than once.
func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return
	}
	q.closed = true
	close(q.events)
}
