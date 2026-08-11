# staging-manual-repo

`ingestd` — a small event ingestion daemon. It accepts events over HTTP,
buffers them in a bounded priority queue, and forwards them upstream through a
retrying, breaker-guarded client.

## Layout

| Path                | Responsibility                                  |
| ------------------- | ----------------------------------------------- |
| `cmd/ingestd`       | Entrypoint, flags, logging, signal handling      |
| `internal/config`   | Environment-driven configuration                 |
| `internal/retry`    | Bounded exponential backoff                      |
| `internal/breaker`  | Circuit breaker around the upstream              |
| `internal/metrics`  | Registry and Prometheus text exposition          |
| `internal/httpapi`  | Health, readiness, and the metrics endpoint      |
| `internal/queue`    | Bounded priority queue                           |
| `internal/worker`   | Worker pool draining the queue                   |

## Running

```sh
go run ./cmd/ingestd -listen :8080 -workers 8 -log-format json
```

| Flag              | Default          | Meaning                                |
| ----------------- | ---------------- | -------------------------------------- |
| `-listen`         | `127.0.0.1:8080` | Admin surface bind address             |
| `-config`         | unset            | Optional config file layered over env  |
| `-log-format`     | `text`           | `text` or `json`                       |
| `-log-level`      | `info`           | `debug`, `info`, `warn`, or `error`    |
| `-workers`        | `4`              | Queue consumers                        |
| `-drain-timeout`  | `15s`            | Time given to in-flight work           |
| `-dry-run`        | `false`          | Validate the invocation and exit       |
| `-version`        | `false`          | Print the build stamp and exit         |

Every problem with an invocation is reported at once, so a mistyped command
does not have to be corrected one flag per run. `SIGINT` and `SIGTERM` both
begin a bounded drain.

## Development

```sh
go build ./...
go test ./...
```
