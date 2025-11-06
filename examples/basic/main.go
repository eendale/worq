package main

import (
	"fmt"
	"time"

	"github.com/eendale/worq/pkg/worker"
)

func main() {
	pool := worker.NewPool(3)
	pool.Start()

	for i := 0; i < 5; i++ {
		n := i
		pool.Enqueue(func() error {
			fmt.Println("Processing job:", n)
			time.Sleep(1 * time.Second)
			return nil
		})
	}

	time.Sleep(3 * time.Second)
	pool.Stop()
}
