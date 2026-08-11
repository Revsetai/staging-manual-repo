// Package worker runs the goroutines that handle queued jobs.
package worker

import (
	"context"
	"log"
	"sync"
)

// Job is a unit of work handed to the pool.
type Job struct {
	ID      string
	Payload []byte
}

// Handler processes a single job.
type Handler func(ctx context.Context, j Job) error

// Pool runs a fixed number of goroutines over a channel of jobs.
//
// TODO: the pool takes a bare channel, so it cannot see the queue's ordering
// or depth. There is no timeout around the handler, no counters, and Stop
// returns as soon as the channel drains whether or not the work finished.
type Pool struct {
	jobs    <-chan Job
	handler Handler
	size    int

	wg sync.WaitGroup

	mu        sync.Mutex
	processed int
	failed    int
}

// NewPool builds a pool over the given job channel.
func NewPool(jobs <-chan Job, h Handler, size int) *Pool {
	if size < 1 {
		size = 4
	}
	return &Pool{jobs: jobs, handler: h, size: size}
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

	for j := range p.jobs {
		err := p.handler(ctx, j)

		p.mu.Lock()
		p.processed++
		if err != nil {
			p.failed++
		}
		p.mu.Unlock()

		if err != nil {
			log.Printf("worker: job %s failed: %v", j.ID, err)
		}
	}
}

// Stop waits for the goroutines to finish.
func (p *Pool) Stop() {
	p.wg.Wait()
}

// Processed reports how many jobs have been handled.
func (p *Pool) Processed() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.processed
}

// Failed reports how many jobs returned an error.
func (p *Pool) Failed() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.failed
}
