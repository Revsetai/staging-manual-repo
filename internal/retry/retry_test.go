package retry

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestDoRetriesUntilItWorks(t *testing.T) {
	calls := 0
	err := Do(context.Background(), 3, func() error {
		calls++
		if calls < 2 {
			return errors.New("nope")
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

func TestDoGivesUp(t *testing.T) {
	err := Do(context.Background(), 2, func() error { return errors.New("nope") })
	if err == nil || !strings.Contains(err.Error(), "2 attempts failed") {
		t.Fatalf("error = %v", err)
	}
}

func TestDoStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Do(ctx, 3, func() error { return errors.New("nope") }); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
}
