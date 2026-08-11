package breaker

import (
	"testing"
	"time"
)

func TestBreakerTripsAfterEnoughFailures(t *testing.T) {
	b := New(3, time.Minute)

	for i := 0; i < 2; i++ {
		b.Failure()
	}
	if !b.Allow() {
		t.Fatal("breaker should still be closed below the threshold")
	}

	b.Failure()
	if b.Allow() {
		t.Fatal("breaker should be open at the threshold")
	}
}

func TestSuccessResetsTheCount(t *testing.T) {
	b := New(2, time.Minute)
	b.Failure()
	b.Success()
	b.Failure()

	if !b.Allow() {
		t.Fatal("a success should have cleared the earlier failure")
	}
}

func TestOpenMirrorsAllow(t *testing.T) {
	b := New(1, time.Minute)
	b.Failure()
	if !b.Open() {
		t.Fatal("Open() should report the tripped breaker")
	}
}
