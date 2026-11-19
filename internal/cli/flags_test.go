package cli

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestParseAppliesDefaults(t *testing.T) {
	opts, err := Parse("ingestd", nil, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.ListenAddr != "127.0.0.1:8080" {
		t.Errorf("listen = %q", opts.ListenAddr)
	}
	if opts.LogFormat != "text" {
		t.Errorf("log format = %q", opts.LogFormat)
	}
	if opts.DrainTimeout != 15*time.Second {
		t.Errorf("drain timeout = %s", opts.DrainTimeout)
	}
}

func TestParseReadsFlags(t *testing.T) {
	opts, err := Parse("ingestd", []string{
		"-listen", "0.0.0.0:7000",
		"-log-format", "json",
		"-drain-timeout", "45s",
		"-dry-run",
	}, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.ListenAddr != "0.0.0.0:7000" {
		t.Errorf("listen = %q", opts.ListenAddr)
	}
	if !opts.DryRun {
		t.Error("dry-run should be set")
	}
	if opts.DrainTimeout != 45*time.Second {
		t.Errorf("drain timeout = %s", opts.DrainTimeout)
	}
}

func TestParseRejectsPositionalArguments(t *testing.T) {
	_, err := Parse("ingestd", []string{"serve"}, io.Discard)
	if err == nil || !strings.Contains(err.Error(), "unexpected argument") {
		t.Fatalf("error = %v, want an unexpected-argument failure", err)
	}
}

func TestParseRejectsUnknownLogFormat(t *testing.T) {
	_, err := Parse("ingestd", []string{"-log-format", "yaml"}, io.Discard)
	if err == nil || !strings.Contains(err.Error(), "log format") {
		t.Fatalf("error = %v, want a log-format failure", err)
	}
}

func TestParseRejectsUnknownLogLevel(t *testing.T) {
	_, err := Parse("ingestd", []string{"-log-level", "trace"}, io.Discard)
	if err == nil || !strings.Contains(err.Error(), "log level") {
		t.Fatalf("error = %v, want a log-level failure", err)
	}
}

func TestParseReadsLogLevel(t *testing.T) {
	opts, err := Parse("ingestd", []string{"-log-level", "debug", "-log-format", "json"}, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.LogLevel != "debug" {
		t.Errorf("log level = %q, want debug", opts.LogLevel)
	}
	if !strings.Contains(opts.String(), "log=json/debug") {
		t.Errorf("String() = %q", opts.String())
	}
}

func TestValidateRequiresPort(t *testing.T) {
	opts := Defaults()
	opts.ListenAddr = "localhost"
	if err := opts.Validate(); err == nil {
		t.Fatal("expected a listen address without a port to be rejected")
	}
}

func TestValidateReportsEveryProblemAtOnce(t *testing.T) {
	opts := Options{ListenAddr: "localhost", LogFormat: "yaml", LogLevel: "loud", Workers: 0}

	err := opts.Validate()
	if err == nil {
		t.Fatal("expected the invalid options to be rejected")
	}
	for _, want := range []string{"must include a port", "log format", "log level", "workers", "drain timeout"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error is missing %q, got %v", want, err)
		}
	}
}

func TestVersionSkipsValidation(t *testing.T) {
	opts, err := Parse("ingestd", []string{"-version", "-workers", "0"}, io.Discard)
	if err != nil {
		t.Fatalf("-version should not be blocked by validation, got %v", err)
	}
	if !opts.ShowVersion {
		t.Error("ShowVersion should be set")
	}
}

func TestParseReadsWorkerCount(t *testing.T) {
	opts, err := Parse("ingestd", []string{"-workers", "12"}, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Workers != 12 {
		t.Errorf("workers = %d, want 12", opts.Workers)
	}
	if !strings.Contains(opts.String(), "workers=12") {
		t.Errorf("String() = %q", opts.String())
	}
}

func TestStringDescribesEnvOnlyConfig(t *testing.T) {
	if got := Defaults().String(); !strings.Contains(got, "<env only>") {
		t.Errorf("String() = %q", got)
	}
}
