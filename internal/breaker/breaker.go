// Package breaker stops calling an upstream that keeps failing.
package breaker

import (
	"errors"
	"sync"
	"time"
)

// ErrOpen is returned when the breaker is refusing calls.
var ErrOpen = errors.New("breaker: open")

// Breaker counts consecutive failures and trips once there are too many.
//
// TODO: consecutive is the wrong measure. An upstream that fails once every
// few minutes never trips this, and one success resets the whole count.
type Breaker struct {
	threshold int
	cooldown  time.Duration

	mu       sync.Mutex
	failures int
	openedAt time.Time
}

// New builds a breaker that trips after threshold failures in a row.
func New(threshold int, cooldown time.Duration) *Breaker {
	if threshold < 1 {
		threshold = 5
	}
	if cooldown <= 0 {
		cooldown = 30 * time.Second
	}
	return &Breaker{threshold: threshold, cooldown: cooldown}
}

// Allow reports whether a call may go ahead.
func (b *Breaker) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.openedAt.IsZero() {
		return true
	}
	if time.Since(b.openedAt) < b.cooldown {
		return false
	}

	b.openedAt = time.Time{}
	b.failures = 0
	return true
}

// Success clears the failure count.
func (b *Breaker) Success() {
	b.mu.Lock()
	b.failures = 0
	b.mu.Unlock()
}

// Failure records a failure and trips the breaker if there are enough.
func (b *Breaker) Failure() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures++
	if b.failures >= b.threshold {
		b.openedAt = time.Now()
	}
}

// Open reports whether the breaker is currently refusing calls.
