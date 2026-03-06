package main

import (
	"log"

	"github.com/hooneun/scorpes/internal/api"
	"github.com/hooneun/scorpes/internal/config"
)

func main() {
	// test
	cfg := config.Load()

	a := api.NewAPI(cfg)

	if err := a.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
