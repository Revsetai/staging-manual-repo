# staging-manual-repo

`ingestd` — a small event ingestion daemon. It accepts events over HTTP,
buffers them in a bounded priority queue, and forwards them upstream through a
retrying, breaker-guarded client.

Nothing in here reaches for a framework. Every dependency is the standard
library, which keeps the failure modes legible: when a scrape is slow or a
shutdown hangs, the cause is in this tree.

## Layout

| Path                | Responsibility                                | Status      |
| ------------------- | --------------------------------------------- | ----------- |
| `internal/config`   | Environment-driven configuration               | done        |
| `internal/retry`    | Bounded exponential backoff                    | in progress |
| `internal/breaker`  | Circuit breaker around the upstream            | in progress |
| `internal/metrics`  | Registry and Prometheus text exposition        | in progress |
| `internal/httpapi`  | Health, readiness, and the metrics endpoint    | in progress |
| `internal/queue`    | Bounded priority queue                         | in progress |
| `internal/worker`   | Worker pool draining the queue                 | in progress |
| `cmd/ingestd`       | Process entrypoint and flag parsing            | in progress |

## Configuration

Every setting is read from the environment on top of a safe local default.
An unset or blank variable falls through to that default; a malformed one is
a startup error naming the variable, not a silently ignored value.

| Variable                 | Default                            | Meaning                        |
| ------------------------ | ---------------------------------- | ------------------------------ |
| `INGESTD_LISTEN_ADDR`    | `127.0.0.1:8080`                   | Admin surface bind address     |
| `INGESTD_ENV`            | `dev`                              | `dev`, `staging`, or `prod`    |
| `INGESTD_LOG_LEVEL`      | `info`                             | `debug`/`info`/`warn`/`error`  |
| `INGESTD_QUEUE_CAPACITY` | `1024`                             | Buffered events                |
| `INGESTD_WORKERS`        | `4`                                | Concurrent consumers           |
| `INGESTD_MAX_BODY_BYTES` | `1048576`                          | Largest accepted request body  |
| `INGESTD_FLUSH_INTERVAL` | `5s`                               | Drain cadence                  |
| `INGESTD_SHUTDOWN_GRACE` | `15s`                              | Time given to in-flight work   |
| `INGESTD_UPSTREAM_URL`   | `http://localhost:9000/v1/events`  | Forwarding target              |
| `INGESTD_UPSTREAM_TOKEN` | unset                              | Required when `INGESTD_ENV=prod` |

## Development

```sh
go build ./...
go test ./...
```
