package metrics

import (
	"strings"
	"sync"
	"testing"
)

func TestCounterIsSharedPerIdentity(t *testing.T) {
	r := NewRegistry()
	a := r.Counter("events_total", Labels{"source": "http"})
	b := r.Counter("events_total", Labels{"source": "http"})
	if a != b {
		t.Fatal("expected the same counter instance for the same identity")
	}

	other := r.Counter("events_total", Labels{"source": "grpc"})
	if a == other {
		t.Fatal("different labels must produce different counters")
	}
}

func TestCounterAddIsConcurrencySafe(t *testing.T) {
	r := NewRegistry()
	c := r.Counter("accepted_total", nil)

	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()

	if got := c.Value(); got != 6400 {
		t.Fatalf("value = %d, want 6400", got)
	}
}

func TestGaugeMovesBothWays(t *testing.T) {
	r := NewRegistry()
	g := r.Gauge("queue_depth", nil)
	g.Set(10)
	g.Add(-4)
	if got := g.Value(); got != 6 {
		t.Fatalf("value = %g, want 6", got)
	}
}

func TestGaugeAddIsConcurrencySafe(t *testing.T) {
	r := NewRegistry()
	g := r.Gauge("queue_depth", nil)

	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				g.Add(1)
				g.Add(-0.5)
			}
		}()
	}
	wg.Wait()

	if got := g.Value(); got != 800 {
		t.Fatalf("value = %g, want 800", got)
	}
}

func TestSnapshotIsSorted(t *testing.T) {
	r := NewRegistry()
	r.Counter("zeta_total", nil).Inc()
	r.Counter("alpha_total", nil).Add(3)
	r.Gauge("mid_gauge", nil).Set(1)

	samples := r.Snapshot()
	if len(samples) != 3 {
		t.Fatalf("len(samples) = %d, want 3", len(samples))
	}
	want := []string{"alpha_total", "mid_gauge", "zeta_total"}
	for i, name := range want {
		if samples[i].Name != name {
			t.Errorf("samples[%d].Name = %q, want %q", i, samples[i].Name, name)
		}
	}
}

func TestSnapshotCarriesTheRenderedIdentity(t *testing.T) {
	r := NewRegistry()
	r.Counter("events_total", Labels{"source": "http"}).Inc()

	samples := r.Snapshot()
	if len(samples) != 1 {
		t.Fatalf("len(samples) = %d, want 1", len(samples))
	}
	want := `events_total{source=http}`
	if samples[0].ID != want {
		t.Errorf("ID = %q, want %q", samples[0].ID, want)
	}
	if !strings.Contains(samples[0].String(), want) {
		t.Errorf("String() = %q", samples[0].String())
	}
}

func TestLabelsKeyIsStable(t *testing.T) {
	first := Labels{"b": "2", "a": "1"}.key()
	second := Labels{"a": "1", "b": "2"}.key()
	if first != second {
		t.Fatalf("label key is not order-independent: %q vs %q", first, second)
	}
}
