package retry

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func fastPolicy(attempts int) Policy {
	return Policy{
		MaxAttempts: attempts,
		BaseDelay:   time.Millisecond,
		MaxDelay:    2 * time.Millisecond,
		Multiplier:  2,
	}
}

func TestDoSucceedsOnSecondAttempt(t *testing.T) {
	calls := 0
	err := Do(context.Background(), fastPolicy(3), func(ctx context.Context, a Attempt) error {
		calls++
		if a.First() {
			return errors.New("transient")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if calls != 2 {
		t.Errorf("calls = %d, want 2", calls)
	}
}

func TestAttemptCarriesThePreviousError(t *testing.T) {
	boom := errors.New("transient")
	var seen []error

	_ = Do(context.Background(), fastPolicy(3), func(ctx context.Context, a Attempt) error {
		seen = append(seen, a.LastErr)
		return boom
	})

	if len(seen) != 3 {
		t.Fatalf("attempts = %d, want 3", len(seen))
	}
	if seen[0] != nil {
		t.Errorf("first attempt LastErr = %v, want nil", seen[0])
	}
	for i, err := range seen[1:] {
		if !errors.Is(err, boom) {
			t.Errorf("attempt %d LastErr = %v, want the previous failure", i+1, err)
		}
	}
}

func TestDoStopsOnPermanentError(t *testing.T) {
	sentinel := errors.New("bad request")
	calls := 0
	err := Do(context.Background(), fastPolicy(5), func(ctx context.Context, a Attempt) error {
		calls++
		return Permanent(sentinel)
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want it to wrap the sentinel", err)
	}
	if calls != 1 {
		t.Errorf("calls = %d, want 1", calls)
	}
}

func TestDoReturnsExhausted(t *testing.T) {
	boom := errors.New("still broken")
	err := Do(context.Background(), fastPolicy(3), func(ctx context.Context, a Attempt) error {
		return boom
	})

	if !errors.Is(err, ErrExhausted) {
		t.Fatalf("error = %v, want it to match ErrExhausted", err)
	}
	if !errors.Is(err, boom) {
		t.Fatalf("error = %v, want it to match the underlying failure too", err)
	}

	var exhausted *ExhaustedError
	if !errors.As(err, &exhausted) {
		t.Fatalf("error = %T, want *ExhaustedError", err)
	}
	if exhausted.Attempts != 3 {
		t.Errorf("Attempts = %d, want 3", exhausted.Attempts)
	}
	if exhausted.Elapsed <= 0 {
		t.Errorf("Elapsed = %s, want a positive duration", exhausted.Elapsed)
	}
	if !strings.Contains(err.Error(), "gave up after 3 attempts") {
		t.Errorf("message = %q", err.Error())
	}
}

func TestDoHonoursCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := Do(ctx, DefaultPolicy(), func(ctx context.Context, a Attempt) error {
		t.Fatal("operation should not run under a cancelled context")
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
}

func TestDelayIsCapped(t *testing.T) {
	p := Policy{BaseDelay: time.Second, MaxDelay: 3 * time.Second, Multiplier: 10}
	for _, attempt := range []int{2, 6, 64, 1024} {
		if got := p.Delay(attempt); got > 3*time.Second {
			t.Errorf("Delay(%d) = %s, want it capped at 3s", attempt, got)
		}
	}
}

func TestJitterUsesTheInjectedSource(t *testing.T) {
	p := Policy{
		BaseDelay:  time.Second,
		MaxDelay:   time.Minute,
		Multiplier: 2,
		Jitter:     0.5,
		Rand:       rand.New(rand.NewSource(1)),
	}
	first := p.Delay(3)

	p.Rand = rand.New(rand.NewSource(1))
	if second := p.Delay(3); first != second {
		t.Fatalf("delays differ across identical seeds: %s vs %s", first, second)
	}
}
