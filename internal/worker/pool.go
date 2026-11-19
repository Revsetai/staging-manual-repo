// Package worker runs the goroutines that handle queued jobs.
package worker

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
)

// EventQueue is the queue behavior the pool needs. queue.Queue implements it.
type EventQueue interface {
	Pop(context.Context) (queue.Event, error)
}

var _ EventQueue = (*queue.Queue)(nil)

// Handler processes a single event.
type Handler func(ctx context.Context, event queue.Event) error

// Stats is a point-in-time view of worker activity.
type Stats struct {
	Processed int `json:"processed"`
	Failed    int `json:"failed"`
	Active    int `json:"active"`
}

// Pool runs a fixed number of goroutines over an event queue.
type Pool struct {
	queue   EventQueue
	handler Handler
	size    int

	wg sync.WaitGroup

	mu    sync.Mutex
	stats Stats
}

// NewPool builds a pool over the given event queue.
func NewPool(events EventQueue, h Handler, size int) *Pool {
	if size < 1 {
		size = 4
	}
	return &Pool{queue: events, handler: h, size: size}
}

// Start launches the goroutines.
func (p *Pool) Start(ctx context.Context) {
	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go p.run(ctx)
	}
}

func (p *Pool) run(ctx context.Context) {
	defer p.wg.Done()

	for {
		event, err := p.queue.Pop(ctx)
		if err != nil {
			if !errors.Is(err, queue.ErrClosed) && !errors.Is(err, context.Canceled) {
				log.Printf("worker: queue stopped: %v", err)
			}
			return
		}

		p.mu.Lock()
		p.stats.Active++
		p.mu.Unlock()

		err = p.handler(ctx, event)

		p.mu.Lock()
		p.stats.Active--
		p.stats.Processed++
		if err != nil {
			p.stats.Failed++
		}
		p.mu.Unlock()

		if err != nil {
			log.Printf("worker: event %s failed: %v", event.ID, err)
		}
	}
}

// Stop waits for the goroutines to finish.
func (p *Pool) Stop() {
	p.wg.Wait()
}

// Processed reports how many jobs have been handled.
func (p *Pool) Processed() int {
	return p.Stats().Processed
}

// Failed reports how many jobs returned an error.
func (p *Pool) Failed() int {
	return p.Stats().Failed
}

// Stats returns a consistent copy of all worker counters.
func (p *Pool) Stats() Stats {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.stats
}
