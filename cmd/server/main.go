package main

import (
	"github.com/AlMkin/metricsalert/internal/config"
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
	"log"
)

func main() {
	cfg := config.GetConfig()

	store := storage.NewMemStorage()

	handlers.SetRepository(store)

	srv := server.NewServer()

	log.Printf("Server is starting at %s\n", cfg.Address)
	if err := srv.Run(cfg.Address); err != nil {
		panic(err)
	}
}
