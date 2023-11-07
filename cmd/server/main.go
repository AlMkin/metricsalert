package main

import (
	"flag"
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
	"os"
)

func main() {
	var addrFlag string

	flag.StringVar(&addrFlag, "a", ":8080", "Address to listen on")
	flag.Parse()

	addr := os.Getenv("ADDRESS")
	if addr == "" {
		addr = addrFlag
	}

	store := storage.NewMemStorage()
	handlers.SetRepository(store)

	srv := server.NewServer()
	if err := srv.Run(addr); err != nil {
		panic(err)
	}
}
