# worq — Lightweight Job Queue for Go

`worq` is a small, dependency-light job queue and worker library for Go applications. It provides:

- An in-memory queue implementation for local processing.
- A Redis-backed queue for distributed producers/consumers.
- A simple worker pool for concurrent job execution.
- A minimal retry policy helper for basic retry/backoff behavior.

This repository contains the packages and simple examples used to demonstrate how to enqueue and process jobs.

## Repository layout

- `cmd/worq-cli` — placeholder CLI (coming soon).
- `examples/basic` — in-memory worker pool example.
- `examples/redis` — enqueue a job to a Redis-backed queue.
- `pkg/job` — `Job` and `RetryPolicy` types.
- `pkg/queue` — `Queue` (in-memory) and `RedisQueue` implementations.
- `pkg/worker` — `Worker` and `Pool` (concurrent workers).
- `pkg/storage` — storage interface (placeholder for extensions).
- `internal/logger` — small slog-based logger used across packages.
- `internal/redis_client` — small Redis client helper.
- `tests` — unit tests covering queue, redis client and worker behaviour.

## Quick examples

These examples assume you have Go installed (go 1.23, per `go.mod`).

Run the in-memory worker pool example:

```powershell
# from repository root
go run ./examples/basic
```

Run the Redis example (requires a Redis server listening at localhost:6379):

```powershell
go run ./examples/redis
```

Build the library or use it in your project by importing the module path from `go.mod`:

```powershell
go get github.com/eendale/worq@latest
```

Or add it to your `go.mod`:

```go
require github.com/eendale/worq v0.0.0
```

## Package summaries (short contract)

- pkg/job
  - Types: `Job` (ID, Handler func() error), `RetryPolicy` (MaxRetries, Delay)
  - Purpose: represent a unit of work and policy for retrying failures.
  - Edge cases: handler returning errors should be retried by caller logic using `RetryPolicy`.

- pkg/queue
  - In-memory `Queue`: buffered channel-based queue. Methods: `Enqueue`, `Dequeue`, `Close`.
  - `RedisQueue`: stores jobs as JSON strings in an RPUSH/LPOP list pair.
  - Contract: `Enqueue` returns an error on failure (e.g., queue full or redis error). `Dequeue` returns a job and possibly an error (Redis) or boolean (in-memory).
  - Notes: `RedisQueue` requires a `*redis.Client` (github.com/redis/go-redis/v9).

- pkg/worker
  - `Pool` spins N goroutines, accepts job functions via `Enqueue(func() error)`, and runs them inside `Worker.Run`.
  - `Start` launches workers; `Stop` signals shutdown and waits for workers to finish.
  - The pool generates a random job ID for enqueued functions.

- pkg/storage
  - Placeholder interface for persistence. No concrete implementation provided in this repository.

## Examples explained

- `examples/basic` demonstrates creating a `worker.Pool`, enqueuing simple functions, waiting, and stopping the pool.
- `examples/redis` demonstrates creating a `RedisQueue`, creating a `job.Job` with an ID and handler, and pushing it into Redis.

## Tests

Run the project's tests with:

```powershell
go test ./...
```

The `tests` directory includes unit tests for queue behavior, redis client helper and worker pool helpers.

## Development notes & assumptions

- The project uses `github.com/redis/go-redis/v9` (see `go.mod`).
- The `cmd/worq-cli` folder is a placeholder; the CLI is not implemented beyond a stub.
- The library favors simplicity and small surface area over advanced features (no ack/nack semantics, no visibility timeouts).

## Recommended next steps / improvements

1. Add a robust job persistence format and visibility timeout for Redis (to support workers failing mid-job).
2. Make `RetryPolicy` usage a first-class feature of the worker or queue (automated retries with backoff).
3. Add a proper CLI for enqueuing and inspecting jobs.
4. Add license and contribution guidelines if this repo becomes public-facing.

## License

No license file is included in this repository. Add a LICENSE file if you intend to publish this code.

---

If you'd like, I can also:

- Generate package-level GoDoc comments for public types and functions.
- Add a simple CI workflow that runs `go test ./...` on push.
- Implement a minimal CLI in `cmd/worq-cli` to enqueue jobs from the command line.

What would you like me to do next?
# worq — Lightweight Job Queue for Go

`worq` is a small, dependency-light job queue and worker library for Go applications. It provides:

- An in-memory queue implementation for local processing.
- A Redis-backed queue for distributed producers/consumers.
- A simple worker pool for concurrent job execution.
- A minimal retry policy helper for basic retry/backoff behavior.

This repository contains the packages and simple examples used to demonstrate how to enqueue and process jobs.

## Repository layout

- `cmd/worq-cli` — placeholder CLI (coming soon).
- `examples/basic` — in-memory worker pool example.
- `examples/redis` — enqueue a job to a Redis-backed queue.
- `pkg/job` — `Job` and `RetryPolicy` types.
- `pkg/queue` — `Queue` (in-memory) and `RedisQueue` implementations.
- `pkg/worker` — `Worker` and `Pool` (concurrent workers).
- `pkg/storage` — storage interface (placeholder for extensions).
- `internal/logger` — small slog-based logger used across packages.
- `internal/redis_client` — small Redis client helper.
- `tests` — unit tests covering queue, redis client and worker behaviour.

## Quick examples

These examples assume you have Go installed (go 1.23, per `go.mod`).

Run the in-memory worker pool example:

```powershell
# from repository root
go run ./examples/basic
```

Run the Redis example (requires a Redis server listening at localhost:6379):

```powershell
go run ./examples/redis
```

Build the library or use it in your project by importing the module path from `go.mod`:

```powershell
go get github.com/eendale/worq@latest
```

Or add it to your `go.mod`:

```go
require github.com/eendale/worq v0.0.0
```

## Package summaries (short contract)

- pkg/job
  - Types: `Job` (ID, Handler func() error), `RetryPolicy` (MaxRetries, Delay)
  - Purpose: represent a unit of work and policy for retrying failures.
  - Edge cases: handler returning errors should be retried by caller logic using `RetryPolicy`.

- pkg/queue
  - In-memory `Queue`: buffered channel-based queue. Methods: `Enqueue`, `Dequeue`, `Close`.
  - `RedisQueue`: stores jobs as JSON strings in an RPUSH/LPOP list pair.
  - Contract: `Enqueue` returns an error on failure (e.g., queue full or redis error). `Dequeue` returns a job and possibly an error (Redis) or boolean (in-memory).
  - Notes: `RedisQueue` requires a `*redis.Client` (github.com/redis/go-redis/v9).

- pkg/worker
  - `Pool` spins N goroutines, accepts job functions via `Enqueue(func() error)`, and runs them inside `Worker.Run`.
  - `Start` launches workers; `Stop` signals shutdown and waits for workers to finish.
  - The pool generates a random job ID for enqueued functions.

- pkg/storage
  - Placeholder interface for persistence. No concrete implementation provided in this repository.

## Examples explained

- `examples/basic` demonstrates creating a `worker.Pool`, enqueuing simple functions, waiting, and stopping the pool.
- `examples/redis` demonstrates creating a `RedisQueue`, creating a `job.Job` with an ID and handler, and pushing it into Redis.

## Tests

Run the project's tests with:

```powershell
go test ./...
```

The `tests` directory includes unit tests for queue behavior, redis client helper and worker pool helpers.

## Development notes & assumptions

- The project uses `github.com/redis/go-redis/v9` (see `go.mod`).
- The `cmd/worq-cli` folder is a placeholder; the CLI is not implemented beyond a stub.
- The library favors simplicity and small surface area over advanced features (no ack/nack semantics, no visibility timeouts).

## Recommended next steps / improvements

1. Add a robust job persistence format and visibility timeout for Redis (to support workers failing mid-job).
2. Make `RetryPolicy` usage a first-class feature of the worker or queue (automated retries with backoff).
3. Add a proper CLI for enqueuing and inspecting jobs.
4. Add license and contribution guidelines if this repo becomes public-facing.

## License

No license file is included in this repository. Add a LICENSE file if you intend to publish this code.

---

If you'd like, I can also:

- Generate package-level GoDoc comments for public types and functions.
- Add a simple CI workflow that runs `go test ./...` on push.
- Implement a minimal CLI in `cmd/worq-cli` to enqueue jobs from the command line.

What would you like me to do next?
# ⚙️ workq — Lightweight Job Queue for Go

`workq` is a minimal, dependency-light job queue library for Go.  
It supports **in-memory** and **Redis-backed** queues for distributed workloads.

## ✨ Features
- Simple worker pool
- Redis-based distributed queue
- Job retries with backoff
- Graceful shutdown
- Extensible interfaces

## 📦 Quick Start
```go
package main

import (
    "fmt"
    "time"
    "github.com/endale/workq/pkg/worker"
)

func main() {
    pool := worker.NewPool(3)
    pool.Start()

    for i := 0; i < 5; i++ {
        jobID := i
        pool.Enqueue(func() error {
            fmt.Println("Running job", jobID)
            time.Sleep(1 * time.Second)
            return nil
        })
    }

    pool.Stop()
}
