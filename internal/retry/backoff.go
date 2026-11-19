// Package retry re-runs an operation that failed, with a delay between attempts.
package retry

import (
	"context"
	"fmt"
	"time"
)

// Delay is how long we wait between attempts.
//
// TODO: this is a flat delay. Every caller retrying the same failing upstream
// wakes up at the same moment, which is exactly when it is least useful.
const Delay = 200 * time.Millisecond

// Do runs fn up to attempts times, waiting Delay between tries.
func Do(ctx context.Context, attempts int, fn func() error) error {
	if attempts < 1 {
		attempts = 1
	}

	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		if i == attempts-1 {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(Delay):
		}
	}

	return fmt.Errorf("retry: %d attempts failed: %v", attempts, err)
}
