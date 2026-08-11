package cli

import (
	"testing"
	"time"
)

func TestOptionsZeroValue(t *testing.T) {
	var opts Options
	if opts.ListenAddr != "" {
		t.Errorf("listen addr = %q, want empty", opts.ListenAddr)
	}
	if opts.DrainTimeout != 0 {
		t.Errorf("drain timeout = %s, want 0", opts.DrainTimeout)
	}
}

func TestOptionsCarryValues(t *testing.T) {
	opts := Options{ListenAddr: "0.0.0.0:9000", Workers: 8, DrainTimeout: time.Minute}
	if opts.Workers != 8 {
		t.Errorf("workers = %d, want 8", opts.Workers)
	}
}
