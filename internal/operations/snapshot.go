// Package operations assembles runtime state for the administrative surface.
package operations

import (
	"context"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
	"github.com/Revsetai/staging-manual-repo/internal/worker"
)

// QueueReader is implemented by queues that expose operational state.
type QueueReader interface {
	Snapshot() queue.Stats
}

// WorkerReader is implemented by worker pools that expose operational state.
type WorkerReader interface {
	Stats() worker.Stats
}

// SnapshotProvider supplies the HTTP API with a complete runtime snapshot.
type SnapshotProvider interface {
	Snapshot(context.Context) Snapshot
}

// Snapshot is the frontend contract for the operations dashboard.
type Snapshot struct {
	CapturedAt time.Time    `json:"capturedAt"`
	UptimeMS   int64        `json:"uptimeMs"`
	Queue      queue.Stats  `json:"queue"`
	Workers    worker.Stats `json:"workers"`
	Health     Health       `json:"health"`
}

// Health summarizes whether the runtime can accept and process more work.
type Health struct {
	Status   string `json:"status"`
	Pressure string `json:"pressure"`
}

// RuntimeSnapshotProvider combines queue and worker state.
type RuntimeSnapshotProvider struct {
	queue   QueueReader
	workers WorkerReader
	started time.Time
	now     func() time.Time
}

var _ SnapshotProvider = (*RuntimeSnapshotProvider)(nil)

// NewRuntimeSnapshotProvider builds a provider over the live runtime.
func NewRuntimeSnapshotProvider(events QueueReader, workers WorkerReader) *RuntimeSnapshotProvider {
	return newRuntimeSnapshotProvider(events, workers, time.Now)
}

func newRuntimeSnapshotProvider(events QueueReader, workers WorkerReader, now func() time.Time) *RuntimeSnapshotProvider {
	return &RuntimeSnapshotProvider{
		queue:   events,
		workers: workers,
		started: now(),
		now:     now,
	}
}

// Snapshot reads the queue and workers once and derives their shared health.
func (p *RuntimeSnapshotProvider) Snapshot(context.Context) Snapshot {
	now := p.now()
	queueStats := p.queue.Snapshot()
	workerStats := p.workers.Stats()

	return Snapshot{
		CapturedAt: now,
		UptimeMS:   now.Sub(p.started).Milliseconds(),
		Queue:      queueStats,
		Workers:    workerStats,
		Health:     deriveHealth(queueStats, workerStats),
	}
}

func deriveHealth(queueStats queue.Stats, workerStats worker.Stats) Health {
	pressure := queuePressure(queueStats)
	status := "healthy"
	if queueStats.Closed {
		status = "stopped"
	} else if pressure == "high" || workerStats.Failed > workerStats.Processed/2 {
		status = "degraded"
	}
	return Health{Status: status, Pressure: pressure}
}

func queuePressure(stats queue.Stats) string {
	if stats.Capacity == 0 || stats.Depth == 0 {
		return "idle"
	}
	ratio := float64(stats.Depth) / float64(stats.Capacity)
	switch {
	case ratio >= 0.8:
		return "high"
	case ratio >= 0.5:
		return "medium"
	default:
		return "low"
	}
}
