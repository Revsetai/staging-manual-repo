package config

import (
	"strings"
	"testing"
	"time"
)

func envFrom(pairs map[string]string) Lookup {
	return func(key string) (string, bool) {
		v, ok := pairs[key]
		return v, ok
	}
}

func TestLoadFromAppliesDefaults(t *testing.T) {
	cfg, err := LoadFrom(envFrom(nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ListenAddr != "127.0.0.1:8080" {
		t.Errorf("listen addr = %q, want 127.0.0.1:8080", cfg.ListenAddr)
	}
	if cfg.WorkerCount != 4 {
		t.Errorf("worker count = %d, want 4", cfg.WorkerCount)
	}
	if cfg.FlushInterval != 5*time.Second {
		t.Errorf("flush interval = %s, want 5s", cfg.FlushInterval)
	}
}

func TestLoadFromOverrides(t *testing.T) {
	cfg, err := LoadFrom(envFrom(map[string]string{
		"INGESTD_LISTEN_ADDR":    "0.0.0.0:9999",
		"INGESTD_ENV":            "staging",
		"INGESTD_MAX_BODY_BYTES": "4096",
		"INGESTD_WORKERS":        "16",
		"INGESTD_FLUSH_INTERVAL": "250ms",
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ListenAddr != "0.0.0.0:9999" {
		t.Errorf("listen addr = %q", cfg.ListenAddr)
	}
	if cfg.WorkerCount != 16 {
		t.Errorf("worker count = %d, want 16", cfg.WorkerCount)
	}
	if cfg.FlushInterval != 250*time.Millisecond {
		t.Errorf("flush interval = %s, want 250ms", cfg.FlushInterval)
	}
	if cfg.Environment != EnvStaging {
		t.Errorf("environment = %q, want staging", cfg.Environment)
	}
	if cfg.MaxBodyBytes != 4096 {
		t.Errorf("max body bytes = %d, want 4096", cfg.MaxBodyBytes)
	}
}

func TestLoadFromRejectsBadInteger(t *testing.T) {
	_, err := LoadFrom(envFrom(map[string]string{"INGESTD_WORKERS": "many"}))
	if err == nil {
		t.Fatal("expected an error for a non-integer worker count")
	}
	if !strings.Contains(err.Error(), "INGESTD_WORKERS") {
		t.Errorf("error should name the offending key, got %v", err)
	}
	if !strings.Contains(err.Error(), "expected an integer") {
		t.Errorf("error should explain what was expected, got %v", err)
	}
}

func TestLoadFromRejectsBadDuration(t *testing.T) {
	_, err := LoadFrom(envFrom(map[string]string{"INGESTD_FLUSH_INTERVAL": "soon"}))
	if err == nil {
		t.Fatal("expected an error for a non-duration flush interval")
	}
	if !strings.Contains(err.Error(), "INGESTD_FLUSH_INTERVAL") {
		t.Errorf("error should name the offending key, got %v", err)
	}
}

func TestEveryBindingHasAUniqueKey(t *testing.T) {
	seen := make(map[string]bool)
	for _, b := range bindings() {
		if seen[b.key] {
			t.Errorf("duplicate binding for %s", b.key)
		}
		seen[b.key] = true
	}
	if len(seen) != 10 {
		t.Errorf("bindings = %d, want 10", len(seen))
	}
}

func TestBlankValuesFallThroughToDefaults(t *testing.T) {
	cfg, err := LoadFrom(envFrom(map[string]string{
		"INGESTD_LISTEN_ADDR": "   ",
		"INGESTD_WORKERS":     "",
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ListenAddr != Default().ListenAddr {
		t.Errorf("listen addr = %q, want the default", cfg.ListenAddr)
	}
	if cfg.WorkerCount != Default().WorkerCount {
		t.Errorf("worker count = %d, want the default", cfg.WorkerCount)
	}
}

func TestLoadFromRejectsUnknownEnvironment(t *testing.T) {
	_, err := LoadFrom(envFrom(map[string]string{"INGESTD_ENV": "qa"}))
	if err == nil {
		t.Fatal("expected an unknown tier to be rejected at parse time")
	}
	if !strings.Contains(err.Error(), "dev, staging, prod") {
		t.Errorf("error should list the valid tiers, got %v", err)
	}
}

func TestEnvironmentIsCaseInsensitive(t *testing.T) {
	cfg, err := LoadFrom(envFrom(map[string]string{
		"INGESTD_ENV":            "PROD",
		"INGESTD_UPSTREAM_TOKEN": "secret",
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Environment.IsProduction() {
		t.Errorf("environment = %q, want prod", cfg.Environment)
	}
}

func TestValidateRejectsZeroBodyLimit(t *testing.T) {
	cfg := Default()
	cfg.MaxBodyBytes = 0
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected a zero body limit to be rejected")
	}
}

func TestValidateRequiresTokenInProd(t *testing.T) {
	cfg := Default()
	cfg.Environment = EnvProd
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected prod without a token to be rejected")
	}

	cfg.UpstreamToken = "secret"
	if err := cfg.Validate(); err != nil {
		t.Fatalf("prod with a token should validate, got %v", err)
	}
}

func TestStringRedactsToken(t *testing.T) {
	cfg := Default()
	cfg.UpstreamToken = "super-secret"
	if strings.Contains(cfg.String(), "super-secret") {
		t.Error("String() leaked the upstream token")
	}
}
