package retry

import "time"

// Backoff doubles Delay with each attempt, up to two seconds.
func Backoff(attempt int) time.Duration {
	d := Delay << attempt
	if d > 2*time.Second {
		return 2 * time.Second
	}
	return d
}
