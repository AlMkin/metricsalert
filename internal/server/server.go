package server

import (
	"github.com/AlMkin/metricsalert/internal/handlers"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	return &Server{mux: mux}
}

func (s *Server) Run(port string) error {
	s.mux.HandleFunc("/update/", handlers.UpdateMetricsHandler)
	return http.ListenAndServe(port, s.mux)
}
