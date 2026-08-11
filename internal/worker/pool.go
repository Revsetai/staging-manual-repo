// Package worker runs the pool of goroutines that drain the queue and hand
// events to a downstream handler.
package worker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
)

// Handler processes a single event. Returning an error marks the event failed.
type Handler func(ctx context.Context, e queue.Event) error

// Options configure a Pool.
type Options struct {
	// Size is the number of goroutines draining the queue.
	Size int
	// HandlerTimeout bounds a single Handler invocation.
	HandlerTimeout time.Duration
	// ShutdownGrace bounds how long Stop waits for in-flight work.
	ShutdownGrace time.Duration
	// Logger receives per-event diagnostics.
	Logger *slog.Logger
}

func (o Options) withDefaults() Options {
	if o.Size < 1 {
		o.Size = 4
	}
	if o.HandlerTimeout <= 0 {
		o.HandlerTimeout = 30 * time.Second
	}
	if o.ShutdownGrace <= 0 {
		o.ShutdownGrace = 15 * time.Second
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
	return o
}

// state is the pool's lifecycle position.
type state int32

const (
	stateIdle state = iota
	stateRunning
	stateStopping
	stateStopped
)

// String renders the state for logs.
func (s state) String() string {
	switch s {
	case stateRunning:
		return "running"
	case stateStopping:
		return "stopping"
	case stateStopped:
		return "stopped"
	default:
		return "idle"
	}
}

// Stats is a point-in-time summary of what the pool has done.
type Stats struct {
	Processed uint64
	Failed    uint64
	InFlight  int
	// Abandoned counts events still buffered when the pool stopped.
	Abandoned uint64
	// TotalLatency is the summed handler duration across every processed
	// event; divide by Processed for the mean.
	TotalLatency time.Duration
	// SlowestLatency is the longest single handler call seen so far.
	SlowestLatency time.Duration
}

// MeanLatency reports the average time spent inside the handler.
func (s Stats) MeanLatency() time.Duration {
	if s.Processed == 0 {
		return 0
	}
	return s.TotalLatency / time.Duration(s.Processed)
}

// Pool drains a queue across a fixed number of goroutines.
type Pool struct {
	opts    Options
	queue   *queue.Queue
	handler Handler

	wg     sync.WaitGroup
	cancel context.CancelFunc
	state  atomic.Int32

	mu    sync.Mutex
	stats Stats
}

// ErrAlreadyRunning is returned by Start when the pool is already running.
var ErrAlreadyRunning = errors.New("worker: pool already running")

// NewPool builds a pool. It does not start any goroutines.
func NewPool(q *queue.Queue, h Handler, opts Options) *Pool {
	return &Pool{
		opts:    opts.withDefaults(),
		queue:   q,
		handler: h,
	}
}

// State reports where the pool is in its lifecycle.
func (p *Pool) State() string { return state(p.state.Load()).String() }

// Start launches the pool's goroutines.
func (p *Pool) Start(ctx context.Context) error {
	if !p.state.CompareAndSwap(int32(stateIdle), int32(stateRunning)) {
		return fmt.Errorf("%w (state %s)", ErrAlreadyRunning, p.State())
	}

	runCtx, cancel := context.WithCancel(ctx)
	p.cancel = cancel

	for i := 0; i < p.opts.Size; i++ {
		p.wg.Add(1)
		go p.run(runCtx, i)
	}
	p.opts.Logger.Info("worker pool started", "size", p.opts.Size)
	return nil
}

// run is one worker goroutine's loop.
func (p *Pool) run(ctx context.Context, id int) {
	defer p.wg.Done()

	for {
		e, err := p.queue.Pop(ctx)
		if err != nil {
			if !errors.Is(err, queue.ErrClosed) && ctx.Err() == nil {
				p.opts.Logger.Error("worker stopped on queue error", "worker", id, "error", err)
			}
			return
		}

		p.track(1)
		started := time.Now()
		err = p.dispatch(ctx, e)
		p.track(-1)

		p.record(err, time.Since(started))
		if err != nil {
			p.opts.Logger.Warn("event failed",
				"worker", id, "event", e.ID, "source", e.Source, "error", err)
		}
	}
}

// dispatch runs the handler under the configured timeout.
func (p *Pool) dispatch(ctx context.Context, e queue.Event) error {
	callCtx, cancel := context.WithTimeout(ctx, p.opts.HandlerTimeout)
	defer cancel()

	if err := p.handler(callCtx, e); err != nil {
		return fmt.Errorf("handle %s: %w", e.ID, err)
	}
	return nil
}

func (p *Pool) track(delta int) {
	p.mu.Lock()
	p.stats.InFlight += delta
	p.mu.Unlock()
}

func (p *Pool) record(err error, latency time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.stats.Processed++
	p.stats.TotalLatency += latency
	p.stats.SlowestLatency = max(p.stats.SlowestLatency, latency)
	if err != nil {
		p.stats.Failed++
	}
}

// Stats returns a copy of the pool's counters.
func (p *Pool) Stats() Stats {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.stats
}

// Stop closes the queue and waits for the workers, up to the shutdown grace.
// Calling it on a pool that is not running is a no-op.
func (p *Pool) Stop() error {
	if !p.state.CompareAndSwap(int32(stateRunning), int32(stateStopping)) {
		return nil
	}

	p.queue.Close()

	finished := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(finished)
	}()

	var err error
	select {
	case <-finished:
	case <-time.After(p.opts.ShutdownGrace):
		err = fmt.Errorf("worker: %d events still in flight after %s",
			p.Stats().InFlight, p.opts.ShutdownGrace)
		if p.cancel != nil {
			p.cancel()
		}
		<-finished
	}

	p.state.Store(int32(stateStopped))

	// Anything still buffered was accepted but never handled. Surface it
	// rather than letting it disappear with the process.
	if abandoned := p.queue.DrainRemaining(); len(abandoned) > 0 {
		p.mu.Lock()
		p.stats.Abandoned = uint64(len(abandoned))
		p.mu.Unlock()

		p.opts.Logger.Warn("events abandoned at shutdown",
			"count", len(abandoned), "first", abandoned[0].ID)
	}

	final := p.Stats()
	p.opts.Logger.Info("worker pool stopped",
		"processed", final.Processed,
		"failed", final.Failed,
		"abandoned", final.Abandoned,
		"mean_latency", final.MeanLatency(),
		"slowest_latency", final.SlowestLatency)
	return err
}
