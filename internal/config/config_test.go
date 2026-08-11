package config

import (
	"strings"
	"testing"
)

func TestLoadUsesDefaults(t *testing.T) {
	cfg := Load()
	if cfg.ListenAddr != "127.0.0.1:8080" {
		t.Errorf("listen addr = %q", cfg.ListenAddr)
	}
	if cfg.WorkerCount != 4 {
		t.Errorf("worker count = %d, want 4", cfg.WorkerCount)
	}
}

func TestLoadReadsTheEnvironment(t *testing.T) {
	t.Setenv("INGESTD_LISTEN_ADDR", "0.0.0.0:9000")
	t.Setenv("INGESTD_WORKERS", "12")

	cfg := Load()
	if cfg.ListenAddr != "0.0.0.0:9000" {
		t.Errorf("listen addr = %q", cfg.ListenAddr)
	}
	if cfg.WorkerCount != 12 {
		t.Errorf("worker count = %d, want 12", cfg.WorkerCount)
	}
}

func TestCheckRejectsZeroWorkers(t *testing.T) {
	cfg := Load()
	cfg.WorkerCount = 0
	if err := cfg.Check(); err == nil {
		t.Fatal("expected zero workers to be rejected")
	}
}

func TestDescribeMentionsTheListenAddr(t *testing.T) {
	if !strings.Contains(Load().Describe(), "listen=127.0.0.1:8080") {
		t.Error("Describe() should mention the listen address")
	}
}
