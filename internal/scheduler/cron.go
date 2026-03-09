package scheduler

import (
	"log"

	"github.com/hooneun/scorpes/internal/job"
	"github.com/hooneun/scorpes/internal/worker"
	"github.com/robfig/cron/v3"
)

type CronScheduler struct {
	cron     *cron.Cron
	jobQueue worker.JobQueue
}

func NewCronScheduler(queue worker.JobQueue) *CronScheduler {
	c := cron.New(
		// seconds 지원
		cron.WithSeconds(),
		cron.WithChain(
			// 이전 job 끝나기 전에 실행 방지
			cron.SkipIfStillRunning(cron.DefaultLogger),
		),
	)

	return &CronScheduler{
		cron:     c,
		jobQueue: queue,
	}
}

func (s *CronScheduler) Start() {
	// 매 분 마다 실행
	/**
	┌──────── second
	│ ┌────── minute
	│ │ ┌──── hour
	│ │ │ ┌── day
	│ │ │ │ ┌ month
	│ │ │ │ │ ┌ weekday
	│ │ │ │ │ │
	* * * * * *
	*/
	_, err := s.cron.AddFunc("0 */1 * * * *", func() {
		s.jobQueue <- func() {
			job.ExampleJob()
			job.HealthCheck()
			log.Println("cron job executed")
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	s.cron.Start()
}

func (s *CronScheduler) Stop() {
	s.cron.Stop()
}
