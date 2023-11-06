package main

import (
	"github.com/AlMkin/metricsalert/internal/handlers"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/internal/storage"
)

func main() {
	store := storage.NewMemStorage()
	handlers.SetRepository(store)

	srv := server.NewServer()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
