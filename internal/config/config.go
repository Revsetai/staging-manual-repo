// Package config reads ingestd's settings out of the environment.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds every runtime setting for the daemon.
type Config struct {
	ListenAddr    string
	Environment   string
	LogLevel      string
	QueueCapacity int
	WorkerCount   int
	FlushInterval time.Duration
	ShutdownGrace time.Duration
	UpstreamURL   string
	UpstreamToken string
}

// Load reads the environment on top of a set of hardcoded defaults.
//
// TODO: a malformed value is currently ignored and the default is kept, which
// means a typo in a deploy manifest is invisible until something misbehaves.
func Load() Config {
	cfg := Config{
		ListenAddr:    "127.0.0.1:8080",
		Environment:   "dev",
		LogLevel:      "info",
		QueueCapacity: 1024,
		WorkerCount:   4,
		FlushInterval: 5 * time.Second,
		ShutdownGrace: 15 * time.Second,
		UpstreamURL:   "http://localhost:9000/v1/events",
	}

	if v := os.Getenv("INGESTD_LISTEN_ADDR"); v != "" {
		cfg.ListenAddr = v
	}
	if v := os.Getenv("INGESTD_ENV"); v != "" {
		cfg.Environment = v
	}
	if v := os.Getenv("INGESTD_LOG_LEVEL"); v != "" {
		cfg.LogLevel = v
	}
	if v := os.Getenv("INGESTD_UPSTREAM_URL"); v != "" {
		cfg.UpstreamURL = v
	}
	if v := os.Getenv("INGESTD_UPSTREAM_TOKEN"); v != "" {
		cfg.UpstreamToken = v
	}
	if v := os.Getenv("INGESTD_QUEUE_CAPACITY"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.QueueCapacity = n
		}
	}
	if v := os.Getenv("INGESTD_WORKERS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.WorkerCount = n
		}
	}
	if v := os.Getenv("INGESTD_FLUSH_INTERVAL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			cfg.FlushInterval = d
		}
	}
	if v := os.Getenv("INGESTD_SHUTDOWN_GRACE"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			cfg.ShutdownGrace = d
		}
	}

	return cfg
}

// Check runs a couple of sanity checks and returns on the first problem.
func (c Config) Check() error {
	if c.ListenAddr == "" {
		return fmt.Errorf("config: listen address is empty")
	}
	if c.QueueCapacity < 1 {
		return fmt.Errorf("config: queue capacity must be positive")
	}
	if c.WorkerCount < 1 {
		return fmt.Errorf("config: worker count must be positive")
	}
	return nil
}

// Describe renders the config for the startup banner.
//
// TODO: this prints the upstream token in the clear.
func (c Config) Describe() string {
	return fmt.Sprintf("listen=%s env=%s log=%s queue=%d workers=%d upstream=%s token=%s",
		c.ListenAddr, c.Environment, c.LogLevel, c.QueueCapacity,
		c.WorkerCount, c.UpstreamURL, c.UpstreamToken)
}
