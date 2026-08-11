// Package httpapi serves the daemon's operational endpoints.
package httpapi

import (
	"fmt"
	"net/http"
	"time"
)

// Server holds the handlers for the admin surface.
//
// TODO: there is no readiness endpoint and no metrics endpoint, so a deploy
// has no way to tell whether this process is ready for traffic.
type Server struct {
	started time.Time
}

// New builds a Server.
func New() *Server {
	return &Server{started: time.Now()}
}

// Handler returns the admin mux.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.handleHealth)
	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "ok uptime=%s\n", time.Since(s.started).Round(time.Second))
}

// Uptime reports how long the server has been up.
func (s *Server) Uptime() time.Duration { return time.Since(s.started) }
