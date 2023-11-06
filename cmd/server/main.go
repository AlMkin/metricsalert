package main

import (
	"flag"
	"github.com/AlMkin/metricsalert/internal/config"
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
	"log"
)

func main() {
	var addrFlag string
	flag.StringVar(&addrFlag, "a", "", "Address to listen on (overrides ADDRESS environment variable)")
	flag.Parse()

	cfg := config.GetServerConfig(addrFlag)

	store := storage.NewMemStorage()
	srv := server.NewServer()

	handlers.SetRepository(store)

	log.Printf("Server is starting at %s\n", cfg.Address)
	if err := srv.Run(cfg.Address); err != nil {
		log.Fatalf("Error when running server: %s", err)
	}
}
