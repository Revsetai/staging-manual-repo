// Command ingestd is the event ingestion daemon.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/cli"
	"github.com/Revsetai/staging-manual-repo/internal/config"
	"github.com/Revsetai/staging-manual-repo/internal/httpapi"
	"github.com/Revsetai/staging-manual-repo/internal/operations"
	"github.com/Revsetai/staging-manual-repo/internal/queue"
	"github.com/Revsetai/staging-manual-repo/internal/worker"
)

func main() {
	opts := cli.Parse()
	cfg := config.Load()
	cfg.ListenAddr = opts.ListenAddr
	cfg.WorkerCount = opts.Workers
	cfg.ShutdownGrace = opts.DrainTimeout
	if err := cfg.Check(); err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, cfg config.Config) error {
	events := queue.New(cfg.QueueCapacity)
	pool := worker.NewPool(events, processEvent, cfg.WorkerCount)
	snapshots := operations.NewRuntimeSnapshotProvider(events, pool)
	api := httpapi.New(events, snapshots)
	server := &http.Server{Addr: cfg.ListenAddr, Handler: api.Handler()}

	pool.Start(ctx)
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("ingestd listening on http://%s with %d workers", cfg.ListenAddr, cfg.WorkerCount)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Printf("shutting down")
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			events.Close()
			pool.Stop()
			return fmt.Errorf("serve operations API: %w", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownGrace)
	defer cancel()
	shutdownErr := server.Shutdown(shutdownCtx)
	events.Close()
	pool.Stop()
	return shutdownErr
}

func processEvent(ctx context.Context, event queue.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(125 * time.Millisecond):
		log.Printf("processed event id=%s source=%s bytes=%d", event.ID, event.Source, len(event.Payload))
		return nil
	}
}
