package router

import (
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Configurator struct {
	router *chi.Mux
}

func NewRouterConfigurator() *Configurator {
	return &Configurator{
		router: chi.NewRouter(),
	}
}

func (rc *Configurator) SetupRoutes() *chi.Mux {
	repo := storage.NewMemStorage()
	handler := handlers.NewHandler(repo)

	rc.router.Post("/update/{type}/{name}/{value}", handler.UpdateMetricsHandler)
	rc.router.Get("/value/{type}/{name}", handler.GetMetricsHandler)
	rc.router.Get("/", handler.ListMetricsHandler)

	return rc.router
}
