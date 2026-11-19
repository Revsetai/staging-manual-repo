# staging-manual-repo

`ingestd` — a small event ingestion daemon used to exercise stacked review
tooling. It accepts events over HTTP, buffers them in a bounded priority
queue, and forwards them upstream through a retrying, breaker-guarded client.

## Layout

| Path                | Responsibility                                  |
| ------------------- | ----------------------------------------------- |
| `cmd/ingestd`       | Process entrypoint and flag parsing              |
| `internal/config`   | Environment-driven configuration                 |
| `internal/retry`    | Bounded exponential backoff                      |
| `internal/breaker`  | Circuit breaker around the upstream              |
| `internal/metrics`  | Registry and Prometheus text exposition          |
| `internal/httpapi`  | Health, readiness, and the metrics endpoint      |
| `internal/queue`    | Bounded priority queue                           |
| `internal/worker`   | Worker pool draining the queue                   |

## Development

```sh
go build ./...
go test ./...
```

## Pipeline

![Event pipeline](docs/pipeline.png)
