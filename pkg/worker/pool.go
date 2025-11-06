package worker

import (
	"math/rand"
	"sync"
	"time"

	"github.com/eendale/worq/pkg/job"
)

type Pool struct {
	numWorkers int
	jobs       chan job.Job
	wg         sync.WaitGroup
	stop       chan struct{}
}

func NewPool(numWorkers int) *Pool {
	return &Pool{
		numWorkers: numWorkers,
		jobs:       make(chan job.Job, 100),
		stop:       make(chan struct{}),
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.numWorkers; i++ {
		w := NewWorker(i + 1)
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case j := <-p.jobs:
					w.Run(j)
				case <-p.stop:
					return
				}
			}
		}()
	}
}

func (p *Pool) Enqueue(fn func() error) {
	j := job.Job{
		ID:      randomID(),
		Handler: fn,
	}
	p.jobs <- j
}

func (p *Pool) Stop() {
	close(p.stop)
	p.wg.Wait()
}

func randomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
