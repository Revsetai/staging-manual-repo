package operations

import (
	"context"
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/queue"
	"github.com/Revsetai/staging-manual-repo/internal/worker"
)

type queueStub struct{ stats queue.Stats }

func (s queueStub) Snapshot() queue.Stats { return s.stats }

type workerStub struct{ stats worker.Stats }

func (s workerStub) Stats() worker.Stats { return s.stats }

func TestRuntimeSnapshotProviderCombinesRuntimeState(t *testing.T) {
	clock := []time.Time{
		time.Unix(100, 0),
		time.Unix(105, 250_000_000),
	}
	provider := newRuntimeSnapshotProvider(
		queueStub{stats: queue.Stats{Depth: 8, Capacity: 10}},
		workerStub{stats: worker.Stats{Processed: 12, Failed: 1, Active: 2}},
		func() time.Time {
			now := clock[0]
			clock = clock[1:]
			return now
		},
	)

	got := provider.Snapshot(context.Background())

	if got.UptimeMS != 5250 {
		t.Fatalf("uptime = %dms, want 5250ms", got.UptimeMS)
	}
	if got.Health.Status != "degraded" || got.Health.Pressure != "high" {
		t.Fatalf("health = %#v, want degraded/high", got.Health)
	}
	if got.Workers.Active != 2 {
		t.Fatalf("active workers = %d, want 2", got.Workers.Active)
	}
}

func TestRuntimeSnapshotProviderReportsStoppedQueue(t *testing.T) {
	provider := newRuntimeSnapshotProvider(
		queueStub{stats: queue.Stats{Capacity: 10, Closed: true}},
		workerStub{},
		time.Now,
	)

	got := provider.Snapshot(context.Background())

	if got.Health.Status != "stopped" {
		t.Fatalf("status = %q, want stopped", got.Health.Status)
	}
}
