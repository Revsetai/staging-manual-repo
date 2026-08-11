// Package queue implements the bounded, priority-ordered buffer that sits
// between ingestd's HTTP intake and its worker pool.
package queue

import (
	"container/heap"
	"context"
	"errors"
	"sync"
	"time"
)

// Priority orders events within the queue. Higher values drain first.
type Priority int

const (
	// PriorityLow is used for backfills and replays.
	PriorityLow Priority = iota
	// PriorityNormal is the default for live traffic.
	PriorityNormal
	// PriorityHigh is reserved for control-plane events.
	PriorityHigh
)

// String renders the priority for logs and metric labels.
func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityNormal:
		return "normal"
	case PriorityHigh:
		return "high"
	default:
		return "unknown"
	}
}

// Event is a single unit of work moving through the pipeline.
type Event struct {
	// ID uniquely identifies the event for deduplication.
	ID string
	// Source names the transport the event arrived on.
	Source string
	// Payload is the opaque body forwarded upstream.
	Payload []byte
	// Priority controls drain order.
	Priority Priority
	// EnqueuedAt is stamped when the event is accepted.
	EnqueuedAt time.Time
}

// Age reports how long the event has been waiting.
func (e Event) Age(now time.Time) time.Duration { return now.Sub(e.EnqueuedAt) }

var (
	// ErrFull is returned when the queue is at capacity.
	ErrFull = errors.New("queue: at capacity")
	// ErrClosed is returned once the queue has been closed.
	ErrClosed = errors.New("queue: closed")
)

// item is the heap's internal record; sequence breaks priority ties in FIFO
// order so equal-priority events keep their arrival ordering.
type item struct {
	event    Event
	sequence uint64
}

type itemHeap []item

func (h itemHeap) Len() int { return len(h) }

func (h itemHeap) Less(i, j int) bool {
	if h[i].event.Priority != h[j].event.Priority {
		return h[i].event.Priority > h[j].event.Priority
	}
	return h[i].sequence < h[j].sequence
}

func (h itemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *itemHeap) Push(x any) { *h = append(*h, x.(item)) }

func (h *itemHeap) Pop() any {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}

// Queue is a bounded priority queue safe for concurrent use.
//
// Waiting consumers are woken over a channel rather than a sync.Cond: a Cond
// cannot participate in a select, so every blocked Pop had to be woken by a
// context watchdog goroutine just to re-check cancellation. The channel folds
// both wake-ups into one select.
type Queue struct {
	capacity int

	mu       sync.Mutex
	items    itemHeap
	sequence uint64
	evicted  uint64
	closed   bool

	// ready receives one token per pushed event.
	ready chan struct{}
	// done is closed exactly once, by Close.
	done chan struct{}
}

// New builds a queue holding at most capacity events.
func New(capacity int) *Queue {
	capacity = max(capacity, 1)
	return &Queue{
		capacity: capacity,
		ready:    make(chan struct{}, capacity),
		done:     make(chan struct{}),
	}
}

// Push enqueues an event.
//
// A full queue rejects the event with ErrFull, unless the arrival outranks
// something already buffered: dropping the least important queued event is
// better than shedding a control-plane message behind a backlog of replays.
// The number of evictions is counted so the caller can alert on it.
func (q *Queue) Push(e Event) error {
	q.mu.Lock()

	if q.closed {
		q.mu.Unlock()
		return ErrClosed
	}
	if len(q.items) >= q.capacity && !q.evictFor(e) {
		q.mu.Unlock()
		return ErrFull
	}
	if e.EnqueuedAt.IsZero() {
		e.EnqueuedAt = time.Now()
	}

	heap.Push(&q.items, item{event: e, sequence: q.sequence})
	q.sequence++
	q.mu.Unlock()

	// A dropped token is harmless: the buffer is already holding at least
	// one wake-up for every event a consumer could claim, and every consumer
	// re-checks the heap before it blocks again.
	select {
	case q.ready <- struct{}{}:
	default:
	}
	return nil
}

// Pop blocks until an event is available, the queue is closed, or ctx is done.
func (q *Queue) Pop(ctx context.Context) (Event, error) {
	for {
		// Checked before the buffer is served: a cancelled consumer must
		// stop taking work, otherwise a hard shutdown keeps handing events
		// to workers that are already being torn down.
		if err := ctx.Err(); err != nil {
			return Event{}, err
		}
		if e, ok := q.TryPop(); ok {
			return e, nil
		}

		select {
		case <-q.ready:
			// A push landed; loop round and claim it. Another consumer may
			// have taken it first, in which case TryPop misses and we wait
			// again.
			if e, ok := q.TryPop(); ok {
				return e, nil
			}
		case <-q.done:
			if e, ok := q.TryPop(); ok {
				return e, nil
			}
			return Event{}, ErrClosed
		case <-ctx.Done():
			return Event{}, ctx.Err()
		}
	}
}

// evictFor makes room for e by discarding the lowest-priority, newest event
// already queued, and reports whether it managed to. The caller must hold
// q.mu and must already know the queue is full.
func (q *Queue) evictFor(e Event) bool {
	victim := -1
	for i := range q.items {
		if q.items[i].event.Priority >= e.Priority {
			continue
		}
		if victim == -1 || less(q.items[victim], q.items[i]) {
			victim = i
		}
	}
	if victim == -1 {
		return false
	}

	q.items[victim] = q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	heap.Init(&q.items)
	q.evicted++
	return true
}

// less reports whether a is a better eviction victim than b: lower priority
// first, then the most recently enqueued.
func less(a, b item) bool {
	if a.event.Priority != b.event.Priority {
		return a.event.Priority > b.event.Priority
	}
	return a.sequence < b.sequence
}

// Evicted reports how many events have been dropped to make room.
func (q *Queue) Evicted() uint64 {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.evicted
}

// TryPop removes the highest-priority event without blocking.
func (q *Queue) TryPop() (Event, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return Event{}, false
	}
	return heap.Pop(&q.items).(item).event, true
}

// DrainRemaining removes and returns every buffered event in drain order.
//
// Shutdown needs the leftovers in one go: popping them one at a time races
// against workers that are still finishing, and the caller ends up holding a
// partial batch it cannot account for.
func (q *Queue) DrainRemaining() []Event {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}
	drained := make([]Event, 0, len(q.items))
	for len(q.items) > 0 {
		drained = append(drained, heap.Pop(&q.items).(item).event)
	}
	return drained
}

// Len reports the number of buffered events.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// Cap reports the queue's capacity.
func (q *Queue) Cap() int { return q.capacity }

// Close wakes every blocked consumer. Buffered events remain drainable
// through TryPop, and Close is safe to call more than once.
func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return
	}
	q.closed = true
	close(q.done)
}

// Closed reports whether Close has been called.
func (q *Queue) Closed() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.closed
}
