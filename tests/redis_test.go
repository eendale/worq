package tests

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/eendale/worq/pkg/job"
	"github.com/eendale/worq/pkg/queue"
)

func TestRedisQueue(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	rq := queue.NewRedisQueue(client, "test_jobs")

	j := job.Job{ID: "redis-job", Handler: func() error { return nil }}

	if err := rq.Enqueue(j); err != nil {
		t.Fatal(err)
	}

	deq, err := rq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}

	if deq.ID != j.ID {
		t.Fatalf("expected %s, got %s", j.ID, deq.ID)
	}
}
