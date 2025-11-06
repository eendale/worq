package queue

import (
	"errors"
	"github.com/eendale/worq/pkg/job"
)

type Queue struct {
	jobs chan job.Job
}

func NewQueue(size int) *Queue {
	return &Queue{
		jobs: make(chan job.Job, size),
	}
}

func (q *Queue) Enqueue(j job.Job) error {
	select {
	case q.jobs <- j:
		return nil
	default:
		return errors.New("queue full")
	}
}

func (q *Queue) Dequeue() (job.Job, bool) {
	j, ok := <-q.jobs
	return j, ok
}

func (q *Queue) Close() {
	close(q.jobs)
}
