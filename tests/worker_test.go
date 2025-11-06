package tests

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/eendale/worq/pkg/worker"
)

func TestWorkerPool(t *testing.T) {
	var counter int32

	pool := worker.NewPool(2)
	pool.Start()

	for i := 0; i < 5; i++ {
		pool.Enqueue(func() error {
			atomic.AddInt32(&counter, 1)
			return nil
		})
	}

	time.Sleep(1 * time.Second)
	pool.Stop()

	if counter != 5 {
		t.Fatalf("expected 5 jobs processed, got %d", counter)
	}
}
