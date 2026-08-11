// Package retry implements a bounded exponential backoff policy used by every
// outbound call ingestd makes.
package retry

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Policy describes how a failing operation is retried.
type Policy struct {
	// MaxAttempts includes the initial attempt. A value below 1 means 1.
	MaxAttempts int
	// BaseDelay is the delay after the first failure.
	BaseDelay time.Duration
	// MaxDelay caps any single computed delay.
	MaxDelay time.Duration
	// Multiplier scales the delay between successive attempts.
	Multiplier float64
	// Jitter is the fraction of the delay randomised away, in [0, 1].
	Jitter float64
	// Rand supplies the jitter. Leave it nil to use the global source; tests
	// set it so that a delay sequence is reproducible.
	Rand *rand.Rand
	// Retryable decides whether an error is worth another attempt. Leave it
	// nil to retry everything that is not explicitly Permanent.
	//
	// A caller that wraps a dependency in something with its own opinion —
	// a circuit breaker, a rate limiter — cannot mark those errors permanent
	// at the point they are created, because they are only meaningless to
	// retry from the perspective of the policy that owns them.
	Retryable func(error) bool
}

// shouldRetry applies the policy's filter on top of the permanent marker.
func (p Policy) shouldRetry(err error) bool {
	if err == nil || IsPermanent(err) {
		return false
	}
	if p.Retryable != nil {
		return p.Retryable(err)
	}
	return true
}

// DefaultPolicy is a conservative policy suited to upstream HTTP calls.
func DefaultPolicy() Policy {
	return Policy{
		MaxAttempts: 5,
		BaseDelay:   100 * time.Millisecond,
		MaxDelay:    10 * time.Second,
		Multiplier:  2.0,
		Jitter:      0.2,
	}
}

// ErrExhausted is returned when every attempt failed.
var ErrExhausted = errors.New("retry: attempts exhausted")

// ExhaustedError reports how much effort was spent before giving up. It is
// returned instead of a bare fmt.Errorf string so that callers can put the
// attempt count and elapsed time onto a metric rather than into a log line
// nobody parses.
type ExhaustedError struct {
	// Attempts is how many tries were made.
	Attempts int
	// Elapsed is the total time spent, including backoff.
	Elapsed time.Duration
	// Last is the final failure.
	Last error
}

func (e *ExhaustedError) Error() string {
	return fmt.Sprintf("retry: gave up after %d attempts in %s: %v",
		e.Attempts, e.Elapsed.Round(time.Millisecond), e.Last)
}

// Unwrap exposes both the sentinel and the underlying failure, so that
// errors.Is finds ErrExhausted and the caller's own sentinels alike.
func (e *ExhaustedError) Unwrap() []error { return []error{ErrExhausted, e.Last} }

// Attempt describes the try an Operation is about to make. Handing the
// operation the previous failure lets callers log or branch on it without
// keeping their own counters.
type Attempt struct {
	// Number is the zero-indexed attempt about to run.
	Number int
	// Elapsed is the time spent since the first attempt started.
	Elapsed time.Duration
	// LastErr is the error from the previous attempt, nil on the first.
	LastErr error
}

// First reports whether this is the initial try.
func (a Attempt) First() bool { return a.Number == 0 }

// Operation is a unit of work that may be retried.
type Operation func(ctx context.Context, attempt Attempt) error

// permanent wraps an error that must not be retried.
type permanent struct{ err error }

func (p permanent) Error() string { return "permanent: " + p.err.Error() }
func (p permanent) Unwrap() error { return p.err }

// Permanent marks err as non-retryable.
func Permanent(err error) error {
	if err == nil {
		return nil
	}
	return permanent{err: err}
}

// IsPermanent reports whether err was marked non-retryable.
func IsPermanent(err error) bool {
	var p permanent
	return errors.As(err, &p)
}

// Delay computes the backoff before the given zero-indexed attempt.
//
// The growth is accumulated iteratively rather than through math.Pow: the
// cap applies at every step, so a large attempt number cannot overflow the
// float before it is clamped.
func (p Policy) Delay(attempt int) time.Duration {
	if attempt <= 0 {
		return 0
	}
	multiplier := p.Multiplier
	if multiplier <= 1 {
		multiplier = 2
	}

	raw := float64(p.BaseDelay)
	for i := 1; i < attempt; i++ {
		raw *= multiplier
		if p.MaxDelay > 0 && raw >= float64(p.MaxDelay) {
			raw = float64(p.MaxDelay)
			break
		}
	}
	if p.MaxDelay > 0 && raw > float64(p.MaxDelay) {
		raw = float64(p.MaxDelay)
	}

	if jitter := min(max(p.Jitter, 0), 1); jitter > 0 {
		spread := raw * jitter
		raw = raw - spread/2 + p.random()*spread
	}
	return time.Duration(max(raw, 0))
}

// random draws from the policy's source, falling back to the global one.
func (p Policy) random() float64 {
	if p.Rand != nil {
		return p.Rand.Float64()
	}
	return rand.Float64()
}

// Do runs op until it succeeds, the policy is exhausted, the error is marked
// permanent, or ctx is cancelled.
func Do(ctx context.Context, p Policy, op Operation) error {
	attempts := max(p.MaxAttempts, 1)
	started := time.Now()

	var last error
	for number := 0; number < attempts; number++ {
		if err := ctx.Err(); err != nil {
			return err
		}

		last = op(ctx, Attempt{
			Number:  number,
			Elapsed: time.Since(started),
			LastErr: last,
		})
		switch {
		case last == nil:
			return nil
		case !p.shouldRetry(last):
			return last
		case number == attempts-1:
			return &ExhaustedError{
				Attempts: attempts,
				Elapsed:  time.Since(started),
				Last:     last,
			}
		}

		timer := time.NewTimer(p.Delay(number + 1))
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}

	// Unreachable: the loop returns on the final attempt.
	return &ExhaustedError{Attempts: attempts, Elapsed: time.Since(started), Last: last}
}
