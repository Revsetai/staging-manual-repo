package metrics

import (
	"strings"
	"testing"
)

func TestEncodeRegistryRendersFamilies(t *testing.T) {
	r := NewRegistry()
	r.Counter("events_total", Labels{"source": "http"}).Add(7)
	r.Counter("events_total", Labels{"source": "grpc"}).Add(2)
	r.Gauge("queue_depth", nil).Set(3)

	enc := NewEncoder("ingestd")
	enc.Describe("events_total", "Events accepted by transport.")

	var out strings.Builder
	if err := enc.EncodeRegistry(&out, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := out.String()

	for _, want := range []string{
		"# HELP ingestd_events_total Events accepted by transport.",
		"# TYPE ingestd_events_total counter",
		`ingestd_events_total{source="grpc"} 2`,
		`ingestd_events_total{source="http"} 7`,
		"# TYPE ingestd_queue_depth gauge",
		"ingestd_queue_depth 3",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output is missing %q\n---\n%s", want, got)
		}
	}
}

func TestEncodeEmitsOneTypeLinePerFamily(t *testing.T) {
	r := NewRegistry()
	for _, source := range []string{"a", "b", "c"} {
		r.Counter("events_total", Labels{"source": source}).Inc()
	}

	var out strings.Builder
	if err := NewEncoder("").EncodeRegistry(&out, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n := strings.Count(out.String(), "# TYPE events_total"); n != 1 {
		t.Fatalf("TYPE lines = %d, want 1", n)
	}
}

func TestLabelValuesAreEscaped(t *testing.T) {
	r := NewRegistry()
	r.Counter("errors_total", Labels{"message": `he said "no"`}).Inc()

	var out strings.Builder
	if err := NewEncoder("").EncodeRegistry(&out, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), `message="he said \"no\""`) {
		t.Fatalf("quotes were not escaped: %s", out.String())
	}
}

func TestEncodeGroupsContiguousRuns(t *testing.T) {
	// Encode consumes Snapshot's ordering, so a family's samples arrive as
	// one contiguous run. Feed it directly to pin that contract down.
	samples := []Sample{
		{ID: "a_total", Name: "a_total", Kind: KindCounter, Value: 1},
		{ID: "b_total", Name: "b_total", Kind: KindCounter, Value: 2},
		{ID: "b_total{x=1}", Name: "b_total", Kind: KindCounter, Labels: Labels{"x": "1"}, Value: 3},
		{ID: "c_gauge", Name: "c_gauge", Kind: KindGauge, Value: 4},
	}

	var out strings.Builder
	if err := NewEncoder("ns").Encode(&out, samples); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := out.String()
	if n := strings.Count(got, "# TYPE"); n != 3 {
		t.Fatalf("TYPE lines = %d, want 3\n%s", n, got)
	}
	if !strings.Contains(got, `ns_b_total{x="1"} 3`) {
		t.Errorf("missing the labelled series:\n%s", got)
	}
}

func TestEncodeHandlesNoSamples(t *testing.T) {
	var out strings.Builder
	if err := NewEncoder("ns").Encode(&out, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Len() != 0 {
		t.Fatalf("expected empty output, got %q", out.String())
	}
}

func TestNamespaceIsOptional(t *testing.T) {
	r := NewRegistry()
	r.Gauge("uptime_seconds", nil).Set(12)

	var out strings.Builder
	if err := NewEncoder("").EncodeRegistry(&out, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "uptime_seconds 12") {
		t.Fatalf("unexpected output: %s", out.String())
	}
}
