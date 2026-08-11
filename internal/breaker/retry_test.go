package breaker_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/breaker"
	"github.com/Revsetai/staging-manual-repo/internal/retry"
)

// policy is a fast policy wired to respect the breaker's verdict.
func policy() retry.Policy {
	return retry.Policy{
		MaxAttempts: 5,
		BaseDelay:   time.Millisecond,
		MaxDelay:    2 * time.Millisecond,
		Multiplier:  2,
		Retryable:   breaker.Retryable,
	}
}

func TestRetryStopsOnceTheCircuitOpens(t *testing.T) {
	b := breaker.New(breaker.Settings{
		FailureThreshold: 2,
		Window:           time.Minute,
		Cooldown:         time.Hour,
	})

	calls := 0
	err := retry.Do(context.Background(), policy(), func(ctx context.Context, a retry.Attempt) error {
		return b.Call(ctx, func(context.Context) error {
			calls++
			return errors.New("upstream down")
		})
	})

	if !errors.Is(err, breaker.ErrOpen) {
		t.Fatalf("error = %v, want ErrOpen", err)
	}
	// Two attempts trip the breaker; the third is refused outright and the
	// policy gives up rather than spending its remaining budget.
	if calls != 2 {
		t.Fatalf("upstream calls = %d, want 2", calls)
	}
}

func TestRetryKeepsGoingWhileTheCircuitIsClosed(t *testing.T) {
	b := breaker.New(breaker.Settings{FailureThreshold: 10, Window: time.Minute})

	calls := 0
	err := retry.Do(context.Background(), policy(), func(ctx context.Context, a retry.Attempt) error {
		return b.Call(ctx, func(context.Context) error {
			calls++
			if calls < 3 {
				return errors.New("transient")
			}
			return nil
		})
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if calls != 3 {
		t.Fatalf("upstream calls = %d, want 3", calls)
	}
}

func TestRetryableOnlyRejectsErrOpen(t *testing.T) {
	if breaker.Retryable(breaker.ErrOpen) {
		t.Error("an open circuit should not be retried")
	}
	if !breaker.Retryable(errors.New("connection reset")) {
		t.Error("an ordinary failure should still be retried")
	}
}
