package server

import (
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	router := chi.NewRouter()
	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	s.router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricsHandler)
	s.router.Get("/value/{type}/{name}", handlers.GetMetricsHandler)
	s.router.Get("/", handlers.ListMetricsHandler)
	return http.ListenAndServe(addr, s.router)
}
