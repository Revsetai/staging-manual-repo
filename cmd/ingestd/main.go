// Command ingestd is the event ingestion daemon.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Revsetai/staging-manual-repo/internal/cli"
)

func main() {
	opts := cli.Parse()
	log.Printf("ingestd starting on %s with %d workers", opts.ListenAddr, opts.Workers)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Printf("shutting down")
	os.Exit(0)
}
