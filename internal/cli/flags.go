// Package cli parses ingestd's command line and turns it into a runnable
// invocation.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"slices"
	"strings"
	"time"
)

// Options are the flag-derived settings for a single ingestd invocation.
type Options struct {
	// ListenAddr is the host:port the admin surface binds to.
	ListenAddr string
	// ConfigPath optionally points at a config file to layer over the env.
	ConfigPath string
	// LogFormat is either "text" or "json".
	LogFormat string
	// LogLevel is one of "debug", "info", "warn", "error".
	LogLevel string
	// Workers is the number of queue consumers to run.
	Workers int
	// DrainTimeout bounds graceful shutdown.
	DrainTimeout time.Duration
	// DryRun parses and validates, then exits without serving.
	DryRun bool
	// ShowVersion prints the build stamp and exits.
	ShowVersion bool
}

// Defaults returns the options used when no flags are supplied.
func Defaults() Options {
	return Options{
		ListenAddr:   "127.0.0.1:8080",
		LogFormat:    "text",
		LogLevel:     "info",
		Workers:      4,
		DrainTimeout: 15 * time.Second,
	}
}

var (
	validLogFormats = []string{"text", "json"}
	validLogLevels  = []string{"debug", "info", "warn", "error"}
)

// register declares every flag against fs, writing straight into o.
func (o *Options) register(fs *flag.FlagSet) {
	fs.StringVar(&o.ListenAddr, "listen", o.ListenAddr, "host:port for the admin surface")
	fs.StringVar(&o.ConfigPath, "config", o.ConfigPath, "path to an optional config file")
	fs.StringVar(&o.LogFormat, "log-format", o.LogFormat, "log output format: text or json")
	fs.StringVar(&o.LogLevel, "log-level", o.LogLevel, "minimum log level: debug, info, warn or error")
	fs.IntVar(&o.Workers, "workers", o.Workers, "number of queue consumers")
	fs.DurationVar(&o.DrainTimeout, "drain-timeout", o.DrainTimeout, "how long to wait for in-flight work")
	fs.BoolVar(&o.DryRun, "dry-run", o.DryRun, "validate configuration and exit")
	fs.BoolVar(&o.ShowVersion, "version", o.ShowVersion, "print the build version and exit")
}

// Parse reads argv (excluding the program name) into Options.
func Parse(name string, argv []string, stderr io.Writer) (Options, error) {
	opts := Defaults()

	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(stderr)
	opts.register(fs)
	fs.Usage = func() {
		fmt.Fprintf(stderr, "usage: %s [flags]\n\nflags:\n", name)
		fs.PrintDefaults()
	}

	if err := fs.Parse(argv); err != nil {
		return Options{}, err
	}
	if rest := fs.Args(); len(rest) > 0 {
		return Options{}, fmt.Errorf("unexpected argument %q", rest[0])
	}

	// -version short-circuits before validation so that a broken config
	// still lets an operator ask which binary they are holding.
	if opts.ShowVersion {
		return opts, nil
	}
	if err := opts.Validate(); err != nil {
		return Options{}, err
	}
	return opts, nil
}

// Validate reports every problem with the parsed options at once, so a
// mistyped invocation does not have to be fixed one flag per run.
func (o Options) Validate() error {
	var problems []string

	switch {
	case strings.TrimSpace(o.ListenAddr) == "":
		problems = append(problems, "listen address must not be empty")
	case !strings.Contains(o.ListenAddr, ":"):
		problems = append(problems, fmt.Sprintf("listen address %q must include a port", o.ListenAddr))
	}
	if !slices.Contains(validLogFormats, o.LogFormat) {
		problems = append(problems, fmt.Sprintf("log format %q must be one of %s",
			o.LogFormat, strings.Join(validLogFormats, ", ")))
	}
	if !slices.Contains(validLogLevels, o.LogLevel) {
		problems = append(problems, fmt.Sprintf("log level %q must be one of %s",
			o.LogLevel, strings.Join(validLogLevels, ", ")))
	}
	if o.Workers < 1 {
		problems = append(problems, fmt.Sprintf("workers must be at least 1, got %d", o.Workers))
	}
	if o.DrainTimeout <= 0 {
		problems = append(problems, fmt.Sprintf("drain timeout must be positive, got %s", o.DrainTimeout))
	}

	if len(problems) == 0 {
		return nil
	}
	return errors.New(strings.Join(problems, "; "))
}

// String renders the options for the startup banner.
func (o Options) String() string {
	config := o.ConfigPath
	if config == "" {
		config = "<env only>"
	}
	return fmt.Sprintf("listen=%s config=%s log=%s/%s workers=%d drain=%s dry-run=%t",
		o.ListenAddr, config, o.LogFormat, o.LogLevel, o.Workers, o.DrainTimeout, o.DryRun)
}
