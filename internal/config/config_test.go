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
	if len(seen) != 15 {
		t.Errorf("bindings = %d, want 15", len(seen))
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

	cfg.Upstream.Token = "secret"
	if err := cfg.Validate(); err != nil {
		t.Fatalf("prod with a token should validate, got %v", err)
	}
}

func TestUpstreamBindingsResolve(t *testing.T) {
	cfg, err := LoadFrom(envFrom(map[string]string{
		"INGESTD_UPSTREAM_URL":          "https://events.internal/v2",
		"INGESTD_UPSTREAM_MAX_ATTEMPTS": "9",
		"INGESTD_RETRY_BASE_DELAY":      "50ms",
		"INGESTD_BREAKER_THRESHOLD":     "20",
		"INGESTD_BREAKER_WINDOW":        "1m",
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Upstream.URL != "https://events.internal/v2" {
		t.Errorf("url = %q", cfg.Upstream.URL)
	}
	if cfg.Upstream.MaxAttempts != 9 {
		t.Errorf("max attempts = %d, want 9", cfg.Upstream.MaxAttempts)
	}
	if cfg.Upstream.RetryBaseDelay != 50*time.Millisecond {
		t.Errorf("retry base delay = %s, want 50ms", cfg.Upstream.RetryBaseDelay)
	}
	if cfg.Upstream.BreakerWindow != time.Minute {
		t.Errorf("breaker window = %s, want 1m", cfg.Upstream.BreakerWindow)
	}
	// Untouched members of the group keep their defaults.
	if cfg.Upstream.BreakerCooldown != Default().Upstream.BreakerCooldown {
		t.Errorf("breaker cooldown = %s, want the default", cfg.Upstream.BreakerCooldown)
	}
}

func TestValidateRejectsZeroAttempts(t *testing.T) {
	cfg := Default()
	cfg.Upstream.MaxAttempts = 0
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected a zero attempt budget to be rejected")
	}
}

func TestStringRedactsToken(t *testing.T) {
	cfg := Default()
	cfg.Upstream.Token = "super-secret"
	if strings.Contains(cfg.String(), "super-secret") {
		t.Error("String() leaked the upstream token")
	}
}
