package breaker

import (
	"context"
	"errors"
	"testing"
	"time"
)

type clock struct{ now time.Time }

func (c *clock) Now() time.Time          { return c.now }
func (c *clock) advance(d time.Duration) { c.now = c.now.Add(d) }

func newTestBreaker(t *testing.T) (*Breaker, *clock) {
	t.Helper()
	c := &clock{now: time.Unix(0, 0)}
	b := New(Settings{
		FailureThreshold: 3,
		Window:           10 * time.Second,
		SuccessThreshold: 2,
		Cooldown:         time.Minute,
		Now:              c.Now,
	})
	return b, c
}

func fail(context.Context) error { return errors.New("upstream down") }

func succeed(context.Context) error { return nil }

func TestBreakerTripsAfterThreshold(t *testing.T) {
	b, _ := newTestBreaker(t)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		if err := b.Call(ctx, fail); err == nil {
			t.Fatalf("call %d should have returned the upstream error", i)
		}
	}

	if got := b.State(); got != StateOpen {
		t.Fatalf("state = %s, want open", got)
	}
	if err := b.Call(ctx, succeed); !errors.Is(err, ErrOpen) {
		t.Fatalf("error = %v, want ErrOpen", err)
	}
}

func TestBreakerRecoversThroughHalfOpen(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		_ = b.Call(ctx, fail)
	}
	c.advance(2 * time.Minute)

	if got := b.State(); got != StateHalfOpen {
		t.Fatalf("state = %s, want half-open after cooldown", got)
	}
	if err := b.Call(ctx, succeed); err != nil {
		t.Fatalf("probe should succeed, got %v", err)
	}
	if err := b.Call(ctx, succeed); err != nil {
		t.Fatalf("second probe should succeed, got %v", err)
	}
	if got := b.State(); got != StateClosed {
		t.Fatalf("state = %s, want closed", got)
	}
}

func TestHalfOpenFailureReopens(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		_ = b.Call(ctx, fail)
	}
	c.advance(2 * time.Minute)
	_ = b.Call(ctx, fail)

	if got := b.State(); got != StateOpen {
		t.Fatalf("state = %s, want open again", got)
	}
}

func TestFailuresOutsideTheWindowDoNotTrip(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()

	// Three failures is the threshold, but spread far enough apart that no
	// two of them are ever inside the same ten-second window.
	for i := 0; i < 5; i++ {
		_ = b.Call(ctx, fail)
		c.advance(30 * time.Second)
	}

	if got := b.State(); got != StateClosed {
		t.Fatalf("state = %s, want closed", got)
	}
	if got := b.Snapshot(); got.Failures != 0 {
		t.Errorf("failures in window = %d, want 0", got.Failures)
	}
}

func TestFailuresInsideTheWindowAccumulate(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()

	_ = b.Call(ctx, fail)
	c.advance(time.Second)
	_ = b.Call(ctx, fail)

	if got := b.Snapshot(); got.Failures != 2 {
		t.Fatalf("failures in window = %d, want 2", got.Failures)
	}
	if got := b.State(); got != StateClosed {
		t.Fatalf("state = %s, want closed below the threshold", got)
	}

	c.advance(time.Second)
	_ = b.Call(ctx, fail)
	if got := b.State(); got != StateOpen {
		t.Fatalf("state = %s, want open at the threshold", got)
	}
}

func TestSuccessDoesNotEraseWindowedFailures(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()

	_ = b.Call(ctx, fail)
	_ = b.Call(ctx, succeed)
	c.advance(time.Second)
	_ = b.Call(ctx, fail)

	if got := b.Snapshot(); got.Failures != 2 {
		t.Fatalf("failures in window = %d, want 2; a success must not clear history", got.Failures)
	}
}

func TestSnapshotReportsHowLongTheCircuitHasBeenOpen(t *testing.T) {
	b, c := newTestBreaker(t)
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		_ = b.Call(ctx, fail)
	}

	c.advance(20 * time.Second)
	snap := b.Snapshot()

	if snap.Healthy() {
		t.Error("an open circuit is not healthy")
	}
	if snap.OpenFor != 20*time.Second {
		t.Errorf("open for = %s, want 20s", snap.OpenFor)
	}
}

func TestSnapshotOfAClosedBreakerHasNoOpenDuration(t *testing.T) {
	b, _ := newTestBreaker(t)
	snap := b.Snapshot()

	if !snap.Healthy() {
		t.Error("a fresh breaker should be healthy")
	}
	if snap.OpenFor != 0 {
		t.Errorf("open for = %s, want 0", snap.OpenFor)
	}
}

func TestResetClosesBreaker(t *testing.T) {
	b, _ := newTestBreaker(t)
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		_ = b.Call(ctx, fail)
	}

	b.Reset()
	if got := b.State(); got != StateClosed {
		t.Fatalf("state = %s, want closed after reset", got)
	}
}
