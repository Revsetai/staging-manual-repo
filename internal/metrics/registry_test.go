package metrics

import "testing"

func TestAddAccumulates(t *testing.T) {
	r := NewRegistry()
	r.Inc("events_total", Labels{"source": "http"})
	r.Add("events_total", Labels{"source": "http"}, 4)

	if got := r.Value("events_total", Labels{"source": "http"}); got != 5 {
		t.Fatalf("value = %g, want 5", got)
	}
}

func TestSetReplaces(t *testing.T) {
	r := NewRegistry()
	r.Set("queue_depth", nil, 9)
	r.Set("queue_depth", nil, 3)

	if got := r.Value("queue_depth", nil); got != 3 {
		t.Fatalf("value = %g, want 3", got)
	}
}

func TestSnapshotIsSorted(t *testing.T) {
	r := NewRegistry()
	r.Inc("zeta_total", nil)
	r.Inc("alpha_total", nil)

	samples := r.Snapshot()
	if len(samples) != 2 {
		t.Fatalf("len = %d, want 2", len(samples))
	}
	if samples[0].Name != "alpha_total" {
		t.Errorf("first sample = %q", samples[0].Name)
	}
}

func TestKeyIsOrderIndependent(t *testing.T) {
	if key("m", Labels{"a": "1", "b": "2"}) != key("m", Labels{"b": "2", "a": "1"}) {
		t.Fatal("label ordering changed the key")
	}
}
