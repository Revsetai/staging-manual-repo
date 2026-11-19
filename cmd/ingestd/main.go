// Command ingestd is the event ingestion daemon.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/cli"
)

// version is stamped at build time with -ldflags.
var version = "dev"

func main() {
	if err := run(os.Args[0], os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "ingestd: %v\n", err)
		os.Exit(1)
	}
}

func run(name string, argv []string) error {
	opts, err := cli.Parse(name, argv, os.Stderr)
	if err != nil {
		return err
	}
	if opts.ShowVersion {
		fmt.Fprintf(os.Stdout, "ingestd %s\n", version)
		return nil
	}

	logger := newLogger(opts)
	logger.Info("starting", "version", version, "options", opts.String())
	logger.Debug("resolved invocation", "argv", argv)

	if opts.DryRun {
		logger.Info("dry run complete; configuration is valid")
		return nil
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return serve(ctx, logger, opts)
}

// serve blocks until the context is cancelled, then drains.
func serve(ctx context.Context, logger *slog.Logger, opts cli.Options) error {
	logger.Info("listening", "addr", opts.ListenAddr, "workers", opts.Workers)

	<-ctx.Done()
	logger.Info("shutdown signal received", "grace", opts.DrainTimeout)

	started := time.Now()
	drainCtx, cancel := context.WithTimeout(context.Background(), opts.DrainTimeout)
	defer cancel()

	if err := drain(drainCtx, opts.Workers); err != nil {
		return fmt.Errorf("drain after %s: %w", time.Since(started).Round(time.Millisecond), err)
	}
	logger.Info("shutdown complete", "took", time.Since(started).Round(time.Millisecond))
	return nil
}

// drain waits for each worker to finish its in-flight event. The per-worker
// loop is a placeholder for the pool's own Stop, which lands with the queue.
func drain(ctx context.Context, workers int) error {
	for worker := range workers {
		select {
		case <-time.After(5 * time.Millisecond):
		case <-ctx.Done():
			return fmt.Errorf("worker %d still busy: %w", worker, ctx.Err())
		}
	}
	return nil
}

// newLogger builds the process logger from the parsed options. The level is
// resolved here rather than in the cli package so that the flag stays a plain
// string and nothing outside main depends on log/slog.
func newLogger(opts cli.Options) *slog.Logger {
	handlerOpts := &slog.HandlerOptions{Level: parseLevel(opts.LogLevel)}
	if opts.LogFormat == "json" {
		return slog.New(slog.NewJSONHandler(os.Stderr, handlerOpts))
	}
	return slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
