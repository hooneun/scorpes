package api

import (
	"fmt"
	"net/http"

	"github.com/hooneun/scorpes/internal/config"
)

type API struct {
	cfg *config.Config
}

func NewAPI(cfg *config.Config) *API {
	return &API{
		cfg: cfg,
	}
}

func (a *API) setupRouter() http.Handler {
	r := NewRouter()

	RegisterRoutes(r)

	return r
}

func (a *API) Run() error {
	addr := fmt.Sprintf(":%d", 8090)
	server := &http.Server{
		Addr:    addr,
		Handler: a.setupRouter(),
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
