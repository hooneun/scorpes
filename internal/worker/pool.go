package worker

import "log"

/**
JobQueue 소비
*/

type Pool struct {
	JobQueue JobQueue
	Workers  int
}

func NewPool(workerCount, queueSize int) *Pool {
	return &Pool{
		JobQueue: make(JobQueue, queueSize),
		Workers:  workerCount,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.Workers; i++ {
		go func(id int) {
			for job := range p.JobQueue {
				log.Printf("worker %d execute job\n", id)
				job()
			}
		}(i)
	}
}
