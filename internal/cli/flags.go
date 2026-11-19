// Package cli parses the daemon's command line.
package cli

import (
	"flag"
	"os"
	"time"
)

// Options are the flag-derived settings for one invocation.
type Options struct {
	ListenAddr   string
	Workers      int
	DrainTimeout time.Duration
}

// Parse reads the process arguments.
//
// TODO: this uses the global flag set, so it cannot be tested without
// reaching into package state, and a bad value exits the process from inside
// a library rather than returning an error.
func Parse() Options {
	var opts Options

	flag.StringVar(&opts.ListenAddr, "listen", "127.0.0.1:8080", "host:port for the admin surface")
	flag.IntVar(&opts.Workers, "workers", 4, "number of workers")
	flag.DurationVar(&opts.DrainTimeout, "drain-timeout", 15*time.Second, "shutdown grace")
	flag.Parse()

	if opts.Workers < 1 {
		flag.Usage()
		os.Exit(2)
	}
	return opts
}

func UnParse() Options {
	var opts Options

	flag.StringVar(&opts.ListenAddr, "listen", "127.0.0.1:8080", "host:port for the admin surface")
	flag.IntVar(&opts.Workers, "workers", 4, "number of workers")
	flag.DurationVar(&opts.DrainTimeout, "drain-timeout", 15*time.Second, "shutdown grace")
	flag.Parse()

	if opts.Workers < 1 {
		flag.Usage()
		os.Exit(2)
	}
	return opts
}
