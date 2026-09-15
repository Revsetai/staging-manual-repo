# staging-manual-repo

`ingestd` — a small event ingestion daemon used to exercise stacked review
tooling. It accepts events over HTTP, buffers them in a bounded queue, drains
them with a worker pool, and exposes a live operations dashboard.

## Layout

| Path                | Responsibility                                  |
| ------------------- | ----------------------------------------------- |
| `cmd/ingestd`       | Process entrypoint and flag parsing              |
| `internal/config`   | Environment-driven configuration                 |
| `internal/retry`    | Bounded exponential backoff                      |
| `internal/breaker`  | Circuit breaker around the upstream              |
| `internal/metrics`  | Registry and Prometheus text exposition          |
| `internal/httpapi`  | Event API and embedded operations dashboard      |
| `internal/operations` | Runtime snapshots and derived health          |
| `internal/queue`    | Bounded priority queue                           |
| `internal/worker`   | Worker pool draining the queue                   |

## Development

```sh
go build ./...
go test ./...
node --test internal/httpapi/web/app.test.mjs
```

Start the daemon and open <http://127.0.0.1:8080> to inspect queue pressure,
worker throughput, and inject demo events through the real intake path.
