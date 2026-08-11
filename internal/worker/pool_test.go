package worker

import (
	"context"
	"errors"
	"testing"
)

func TestPoolHandlesEveryJob(t *testing.T) {
	jobs := make(chan Job, 4)
	p := NewPool(jobs, func(context.Context, Job) error { return nil }, 2)
	p.Start(context.Background())

	for _, id := range []string{"a", "b", "c", "d"} {
		jobs <- Job{ID: id}
	}
	close(jobs)
	p.Stop()

	if got := p.Processed(); got != 4 {
		t.Fatalf("processed = %d, want 4", got)
	}
}

func TestPoolCountsFailures(t *testing.T) {
	jobs := make(chan Job, 2)
	p := NewPool(jobs, func(_ context.Context, j Job) error {
		if j.ID == "bad" {
			return errors.New("boom")
		}
		return nil
	}, 1)
	p.Start(context.Background())

	jobs <- Job{ID: "good"}
	jobs <- Job{ID: "bad"}
	close(jobs)
	p.Stop()

	if got := p.Failed(); got != 1 {
		t.Fatalf("failed = %d, want 1", got)
	}
}
