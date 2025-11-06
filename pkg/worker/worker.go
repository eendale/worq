package worker

import (
	"github.com/eendale/worq/internal/logger"
	"github.com/eendale/worq/pkg/job"
)

type Worker struct {
	id int
}

func NewWorker(id int) *Worker {
	return &Worker{id: id}
}

func (w *Worker) Run(j job.Job) {
	logger.Log.Info("worker starting job", "worker_id", w.id, "job_id", j.ID)
	err := j.Handler()
	if err != nil {
		logger.Log.Error("job failed", "worker_id", w.id, "job_id", j.ID, "error", err)
	} else {
		logger.Log.Info("job completed", "worker_id", w.id, "job_id", j.ID)
	}
}
