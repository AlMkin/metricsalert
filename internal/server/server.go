package server

import (
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	s.router.HandleFunc("/update/{type}/{name}/{value}", handlers.UpdateMetricsHandler).Methods(http.MethodPost)
	s.router.HandleFunc("/value/{type}/{name}", handlers.GetMetricsHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/", handlers.ListMetricsHandler).Methods(http.MethodGet)
	return http.ListenAndServe(addr, s.router)
}
