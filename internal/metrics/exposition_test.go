package metrics

import (
	"strings"
	"testing"
)

func TestEncodeWritesEverySample(t *testing.T) {
	samples := []Sample{
		{Name: "events_total", Labels: Labels{"source": "http"}, Value: 7},
		{Name: "queue_depth", Value: 3},
	}

	var out strings.Builder
	if err := Encode(&out, samples); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := out.String()
	if !strings.Contains(got, `events_total{source="http"} 7`) {
		t.Errorf("missing the counter:\n%s", got)
	}
	if !strings.Contains(got, "queue_depth 3") {
		t.Errorf("missing the gauge:\n%s", got)
	}
}

func TestEncodeSortsLabels(t *testing.T) {
	samples := []Sample{{Name: "m", Labels: Labels{"b": "2", "a": "1"}, Value: 1}}

	var out strings.Builder
	if err := Encode(&out, samples); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), `m{a="1",b="2"}`) {
		t.Errorf("labels are not sorted: %s", out.String())
	}
}
