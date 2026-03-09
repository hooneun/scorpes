package main

import (
	"log"

	"github.com/hooneun/scorpes/internal/api"
	"github.com/hooneun/scorpes/internal/config"
	"github.com/hooneun/scorpes/internal/scheduler"
	"github.com/hooneun/scorpes/internal/worker"
)

func main() {
	cfg := config.Load()

	a := api.NewAPI(cfg)

	pool := worker.NewPool(5, 100)
	pool.Start()

	cronScheduler := scheduler.NewCronScheduler(pool.JobQueue)
	cronScheduler.Start()

	if err := a.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
