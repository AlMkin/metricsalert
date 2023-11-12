package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	router *chi.Mux
}

func NewServer(router *chi.Mux) *Server {
	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
