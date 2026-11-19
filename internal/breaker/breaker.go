// Package breaker implements a three-state circuit breaker that sheds load
// when an upstream dependency starts failing.
package breaker

import (
	"context"
	"errors"
	"sync"
	"time"
)

// State is the breaker's current disposition toward the protected dependency.
type State int

const (
	// StateClosed lets every call through.
	StateClosed State = iota
	// StateOpen rejects every call until the cooldown elapses.
	StateOpen
	// StateHalfOpen admits a limited number of probe calls.
	StateHalfOpen
)

// String renders the state for logs and metrics labels.
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// ErrOpen is returned when the breaker refuses a call outright.
var ErrOpen = errors.New("breaker: circuit is open")

// Settings tune when the breaker trips and how it recovers.
type Settings struct {
	// FailureThreshold is the number of failures within Window that trip it.
	FailureThreshold int
	// Window is how far back failures are counted. A dependency that fails
	// once an hour should not trip a breaker just because those failures
	// happen to be consecutive, which is what a bare counter would do.
	Window time.Duration
	// SuccessThreshold is the number of consecutive half-open successes
	// required to close the circuit again.
	SuccessThreshold int
	// Cooldown is how long the breaker stays open before probing.
	Cooldown time.Duration
	// HalfOpenLimit bounds concurrent probe calls.
	HalfOpenLimit int
	// Now is injectable for tests; defaults to time.Now.
	Now func() time.Time
}

// DefaultSettings returns settings appropriate for a chatty HTTP upstream.
func DefaultSettings() Settings {
	return Settings{
		FailureThreshold: 5,
		Window:           10 * time.Second,
		SuccessThreshold: 2,
		Cooldown:         30 * time.Second,
		HalfOpenLimit:    1,
	}
}

// Breaker guards a single dependency. The zero value is not usable; call New.
type Breaker struct {
	settings Settings

	mu           sync.Mutex
	state        State
	failures     []time.Time
	successes    int
	halfOpenBusy int
	openedAt     time.Time
}

// New builds a breaker, filling in any unset setting with its default.
func New(s Settings) *Breaker {
	defaults := DefaultSettings()
	if s.FailureThreshold < 1 {
		s.FailureThreshold = defaults.FailureThreshold
	}
	if s.Window <= 0 {
		s.Window = defaults.Window
	}
	if s.SuccessThreshold < 1 {
		s.SuccessThreshold = defaults.SuccessThreshold
	}
	if s.Cooldown <= 0 {
		s.Cooldown = defaults.Cooldown
	}
	if s.HalfOpenLimit < 1 {
		s.HalfOpenLimit = defaults.HalfOpenLimit
	}
	if s.Now == nil {
		s.Now = time.Now
	}
	return &Breaker{settings: s, state: StateClosed}
}

// State reports the breaker's current state.
func (b *Breaker) State() State {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.refresh()
	return b.state
}

// refresh promotes an open breaker to half-open once the cooldown elapses.
// The caller must hold b.mu.
func (b *Breaker) refresh() {
	if b.state != StateOpen {
		return
	}
	if b.settings.Now().Sub(b.openedAt) < b.settings.Cooldown {
		return
	}
	b.state = StateHalfOpen
	b.successes = 0
	b.halfOpenBusy = 0
	b.failures = b.failures[:0]
}

// expire drops failures that have aged out of the window.
// The caller must hold b.mu.
func (b *Breaker) expire(now time.Time) {
	cutoff := now.Add(-b.settings.Window)
	keep := b.failures[:0]
	for _, at := range b.failures {
		if at.After(cutoff) {
			keep = append(keep, at)
		}
	}
	b.failures = keep
}

// Failures reports how many failures are currently inside the window.
func (b *Breaker) Failures() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.expire(b.settings.Now())
	return len(b.failures)
}

// allow reserves capacity for one call, or explains why it cannot.
func (b *Breaker) allow() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.refresh()

	switch b.state {
	case StateOpen:
		return ErrOpen
	case StateHalfOpen:
		if b.halfOpenBusy >= b.settings.HalfOpenLimit {
			return ErrOpen
		}
		b.halfOpenBusy++
		return nil
	default:
		return nil
	}
}

// record folds the outcome of a call back into the breaker's state.
func (b *Breaker) record(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.state == StateHalfOpen && b.halfOpenBusy > 0 {
		b.halfOpenBusy--
	}

	now := b.settings.Now()
	if err != nil {
		b.expire(now)
		b.failures = append(b.failures, now)
		b.successes = 0

		if b.state == StateHalfOpen || len(b.failures) >= b.settings.FailureThreshold {
			b.state = StateOpen
			b.openedAt = now
			b.failures = b.failures[:0]
		}
		return
	}

	if b.state == StateHalfOpen {
		b.successes++
		if b.successes >= b.settings.SuccessThreshold {
			b.state = StateClosed
			b.successes = 0
		}
	}
}

// Call runs fn under the breaker's protection.
func (b *Breaker) Call(ctx context.Context, fn func(context.Context) error) error {
	if err := b.allow(); err != nil {
		return err
	}
	err := fn(ctx)
	b.record(err)
	return err
}

// Retryable reports whether err is worth another attempt from the breaker's
// point of view. Retrying straight into an open circuit only burns the
// caller's budget on calls the breaker has already decided to reject, so it
// is wired into retry.Policy.Retryable rather than left to each call site.
func Retryable(err error) bool { return !errors.Is(err, ErrOpen) }

// Reset forces the breaker back to closed, discarding all counters.
func (b *Breaker) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.state = StateClosed
	b.failures = nil
	b.successes = 0
	b.halfOpenBusy = 0
	b.openedAt = time.Time{}
}
