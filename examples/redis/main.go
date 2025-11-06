package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/eendale/worq/pkg/job"
	"github.com/eendale/worq/pkg/queue"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	rq := queue.NewRedisQueue(client, "workq_jobs")

	job := job.Job{
		ID: "example-job",
		Handler: func() error {
			fmt.Println("Running job from Redis queue!")
			return nil
		},
	}

	err := rq.Enqueue(job)
	if err != nil {
		log.Fatal("Failed to enqueue:", err)
	}

	fmt.Println("Job enqueued successfully!")
}
