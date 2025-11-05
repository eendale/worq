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
