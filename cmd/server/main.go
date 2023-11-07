package main

import (
	"flag"
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", ":8080", "Address to listen on")
	flag.Parse()

	store := storage.NewMemStorage()
	handlers.SetRepository(store)

	srv := server.NewServer()
	if err := srv.Run(addr); err != nil {
		panic(err)
	}
}
