package tests

import (
	"testing"

	"github.com/eendale/worq/pkg/job"
	"github.com/eendale/worq/pkg/queue"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	q := queue.NewQueue(2)
	j := job.Job{ID: "1", Handler: func() error { return nil }}

	err := q.Enqueue(j)
	if err != nil {
		t.Fatal(err)
	}

	res, ok := q.Dequeue()
	if !ok {
		t.Fatal("expected job, got none")
	}
	if res.ID != j.ID {
		t.Fatalf("expected ID %s, got %s", j.ID, res.ID)
	}
}
