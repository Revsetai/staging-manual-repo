// Package config loads the ingest daemon's runtime configuration from the
// process environment, applying defaults and validating the result.
package config

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Config is the fully resolved runtime configuration for ingestd.
type Config struct {
	// ListenAddr is the host:port the HTTP surface binds to.
	ListenAddr string
	// Environment names the deployment this process belongs to.
	Environment Environment
	// LogLevel is one of "debug", "info", "warn", "error".
	LogLevel string
	// QueueCapacity bounds the number of buffered events.
	QueueCapacity int
	// WorkerCount is the number of concurrent consumers.
	WorkerCount int
	// FlushInterval is how often buffered events are drained.
	FlushInterval time.Duration
	// ShutdownGrace is how long in-flight work gets to finish.
	ShutdownGrace time.Duration
	// MaxBodyBytes caps the size of a single accepted request body.
	MaxBodyBytes int
	// Upstream describes the forwarding target and how failures against it
	// are handled. These settings were flat fields on Config until the
	// retry and breaker packages landed and the group grew past two.
	Upstream Upstream
}

// Upstream is the forwarding target and its failure-handling policy.
type Upstream struct {
	// URL is where accepted events are forwarded.
	URL string
	// Token authenticates against URL.
	Token string
	// MaxAttempts bounds how often a single event is retried.
	MaxAttempts int
	// RetryBaseDelay is the backoff after the first failure.
	RetryBaseDelay time.Duration
	// BreakerThreshold is the number of failures within BreakerWindow that
	// opens the circuit.
	BreakerThreshold int
	// BreakerWindow is how far back those failures are counted.
	BreakerWindow time.Duration
	// BreakerCooldown is how long the circuit stays open before probing.
	BreakerCooldown time.Duration
}

// Default returns the configuration used when nothing is set in the
// environment. It is deliberately safe for local development.
func Default() Config {
	return Config{
		ListenAddr:    "127.0.0.1:8080",
		Environment:   EnvDev,
		LogLevel:      "info",
		QueueCapacity: 1024,
		WorkerCount:   4,
		FlushInterval: 5 * time.Second,
		ShutdownGrace: 15 * time.Second,
		MaxBodyBytes:  1 << 20,
		Upstream: Upstream{
			URL:              "http://localhost:9000/v1/events",
			MaxAttempts:      5,
			RetryBaseDelay:   100 * time.Millisecond,
			BreakerThreshold: 5,
			BreakerWindow:    10 * time.Second,
			BreakerCooldown:  30 * time.Second,
		},
	}
}

// Load reads configuration from the environment on top of Default.
func Load() (Config, error) {
	return LoadFrom(os.LookupEnv)
}

// Lookup mirrors os.LookupEnv so callers can inject a fake environment.
type Lookup func(key string) (string, bool)

// binding ties one environment variable to the field it populates. Keeping
// the table beside the struct means a new setting is one line, not four
// hand-written lookups that quietly drift out of sync.
type binding struct {
	key   string
	apply func(cfg *Config, raw string) error
}

func bindings() []binding {
	return []binding{
		{"INGESTD_LISTEN_ADDR", assign(func(c *Config) *string { return &c.ListenAddr }, parseString)},
		{"INGESTD_ENV", assign(func(c *Config) *Environment { return &c.Environment }, parseEnvironment)},
		{"INGESTD_LOG_LEVEL", assign(func(c *Config) *string { return &c.LogLevel }, parseString)},
		{"INGESTD_UPSTREAM_URL", assign(func(c *Config) *string { return &c.Upstream.URL }, parseString)},
		{"INGESTD_UPSTREAM_TOKEN", assign(func(c *Config) *string { return &c.Upstream.Token }, parseString)},
		{"INGESTD_UPSTREAM_MAX_ATTEMPTS", assign(func(c *Config) *int { return &c.Upstream.MaxAttempts }, parseInt)},
		{"INGESTD_RETRY_BASE_DELAY", assign(func(c *Config) *time.Duration { return &c.Upstream.RetryBaseDelay }, parseDuration)},
		{"INGESTD_BREAKER_THRESHOLD", assign(func(c *Config) *int { return &c.Upstream.BreakerThreshold }, parseInt)},
		{"INGESTD_BREAKER_WINDOW", assign(func(c *Config) *time.Duration { return &c.Upstream.BreakerWindow }, parseDuration)},
		{"INGESTD_BREAKER_COOLDOWN", assign(func(c *Config) *time.Duration { return &c.Upstream.BreakerCooldown }, parseDuration)},
		{"INGESTD_QUEUE_CAPACITY", assign(func(c *Config) *int { return &c.QueueCapacity }, parseInt)},
		{"INGESTD_WORKERS", assign(func(c *Config) *int { return &c.WorkerCount }, parseInt)},
		{"INGESTD_MAX_BODY_BYTES", assign(func(c *Config) *int { return &c.MaxBodyBytes }, parseInt)},
		{"INGESTD_FLUSH_INTERVAL", assign(func(c *Config) *time.Duration { return &c.FlushInterval }, parseDuration)},
		{"INGESTD_SHUTDOWN_GRACE", assign(func(c *Config) *time.Duration { return &c.ShutdownGrace }, parseDuration)},
	}
}

// assign adapts a field selector and a parser into a binding's apply func.
func assign[T any](field func(*Config) *T, parse func(string) (T, error)) func(*Config, string) error {
	return func(cfg *Config, raw string) error {
		value, err := parse(raw)
		if err != nil {
			return err
		}
		*field(cfg) = value
		return nil
	}
}

func parseString(raw string) (string, error) { return raw, nil }

func parseEnvironment(raw string) (Environment, error) {
	env := Environment(strings.ToLower(raw))
	if !slices.Contains(validEnvironments, env) {
		return "", fmt.Errorf("expected one of %s", joinEnvironments())
	}
	return env, nil
}

func parseInt(raw string) (int, error) {
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errors.New("expected an integer")
	}
	return n, nil
}

func parseDuration(raw string) (time.Duration, error) {
	d, err := time.ParseDuration(raw)
	if err != nil {
		return 0, errors.New("expected a duration such as 250ms or 30s")
	}
	return d, nil
}

// LoadFrom resolves configuration using the supplied lookup function.
func LoadFrom(lookup Lookup) (Config, error) {
	cfg := Default()

	for _, b := range bindings() {
		raw, ok := lookup(b.key)
		if !ok {
			continue
		}
		if raw = strings.TrimSpace(raw); raw == "" {
			continue
		}
		if err := b.apply(&cfg, raw); err != nil {
			return Config{}, fmt.Errorf("config: %s: %w", b.key, err)
		}
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// Environment names a deployment tier. It is a distinct type so that a
// mistyped tier is caught where it is parsed rather than at the one call site
// that happens to compare it against a literal.
type Environment string

const (
	EnvDev     Environment = "dev"
	EnvStaging Environment = "staging"
	EnvProd    Environment = "prod"
)

// IsProduction reports whether the tier carries live traffic.
func (e Environment) IsProduction() bool { return e == EnvProd }

// String renders the tier.
func (e Environment) String() string { return string(e) }

var validEnvironments = []Environment{EnvDev, EnvStaging, EnvProd}

func joinEnvironments() string {
	names := make([]string, len(validEnvironments))
	for i, env := range validEnvironments {
		names[i] = env.String()
	}
	return strings.Join(names, ", ")
}

var validLogLevels = []string{"debug", "info", "warn", "error"}

// Validate reports whether the configuration is internally consistent.
func (c Config) Validate() error {
	var problems []string

	if c.ListenAddr == "" {
		problems = append(problems, "listen address must not be empty")
	}
	if !slices.Contains(validEnvironments, c.Environment) {
		problems = append(problems, fmt.Sprintf("environment %q must be one of %s", c.Environment, joinEnvironments()))
	}
	if !slices.Contains(validLogLevels, c.LogLevel) {
		problems = append(problems, fmt.Sprintf("log level %q must be one of %s", c.LogLevel, strings.Join(validLogLevels, ", ")))
	}
	if c.QueueCapacity < 1 {
		problems = append(problems, "queue capacity must be at least 1")
	}
	if c.WorkerCount < 1 {
		problems = append(problems, "worker count must be at least 1")
	}
	if c.FlushInterval <= 0 {
		problems = append(problems, "flush interval must be positive")
	}
	if c.ShutdownGrace <= 0 {
		problems = append(problems, "shutdown grace must be positive")
	}
	if c.MaxBodyBytes < 1 {
		problems = append(problems, "max body bytes must be at least 1")
	}
	if c.Environment.IsProduction() && c.Upstream.Token == "" {
		problems = append(problems, "upstream token is required in prod")
	}
	if c.Upstream.MaxAttempts < 1 {
		problems = append(problems, "upstream max attempts must be at least 1")
	}
	if c.Upstream.BreakerThreshold < 1 {
		problems = append(problems, "breaker threshold must be at least 1")
	}
	if c.Upstream.BreakerWindow <= 0 {
		problems = append(problems, "breaker window must be positive")
	}

	if len(problems) == 0 {
		return nil
	}
	return errors.New("config: " + strings.Join(problems, "; "))
}

// String renders the configuration for logging, redacting the token.
func (c Config) String() string {
	token := "<unset>"
	if c.Upstream.Token != "" {
		token = "<redacted>"
	}
	return fmt.Sprintf(
		"listen=%s env=%s log=%s queue=%d workers=%d max-body=%d flush=%s grace=%s %s token=%s",
		c.ListenAddr, c.Environment, c.LogLevel, c.QueueCapacity, c.WorkerCount,
		c.MaxBodyBytes, c.FlushInterval, c.ShutdownGrace, c.Upstream, token,
	)
}

// String renders the upstream settings for the startup banner. The token is
// deliberately absent: Config.String is responsible for redacting it.
func (u Upstream) String() string {
	return fmt.Sprintf("upstream=%s attempts=%d backoff=%s breaker=%d/%s cooldown=%s",
		u.URL, u.MaxAttempts, u.RetryBaseDelay,
		u.BreakerThreshold, u.BreakerWindow, u.BreakerCooldown)
}
