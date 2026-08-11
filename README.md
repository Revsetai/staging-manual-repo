# staging-manual-repo

`ingestd` — a small event ingestion daemon. It accepts events over HTTP,
buffers them in a bounded priority queue, and forwards them upstream through a
retrying, breaker-guarded client.

## Backpressure

The queue is bounded, so intake has to decide what to do when it fills. An
arriving event of equal or lower priority is rejected outright; a
higher-priority one evicts the least important event already buffered, newest
first, on the grounds that the older event has already waited and dropping it
wastes more work. Evictions are counted rather than logged, because under real
backpressure the log line is the expensive part.

Shutdown is symmetrical: workers finish what they hold, and anything still
buffered when the grace expires is reported as abandoned instead of vanishing
with the process.

## Layout

| Path                | Responsibility                                  |
| ------------------- | ----------------------------------------------- |
| `cmd/ingestd`       | Process entrypoint and flag parsing              |
| `internal/config`   | Environment-driven configuration                 |
| `internal/retry`    | Bounded exponential backoff                      |
| `internal/breaker`  | Circuit breaker around the upstream              |
| `internal/metrics`  | Registry and Prometheus text exposition          |
| `internal/httpapi`  | Health, readiness, and the metrics endpoint      |
| `internal/queue`    | Bounded priority queue with eviction              |
| `internal/worker`   | Worker pool, latency accounting, graceful stop    |

## Development

```sh
go build ./...
go test ./...
```
