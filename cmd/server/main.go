package main

import (
	"flag"
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
	"log"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", ":8080", "Address to listen on")
	flag.Parse()

	store := storage.NewMemStorage()
	handlers.SetRepository(store)

	srv := server.NewServer()
	log.Printf("Server is starting at %s\n", addr)
	if err := srv.Run(addr); err != nil {
		log.Fatalf("Error when running server: %s", err)
	}
}
