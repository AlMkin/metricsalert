package main

import (
	"flag"
	"github.com/AlMkin/metricsalert/internal/router"
	"github.com/AlMkin/metricsalert/internal/server"
	"github.com/AlMkin/metricsalert/pkg/config"
	"log"
)

func main() {
	addrFlag := flag.String("a", ":8080", "Address to listen on")
	flag.Parse()

	addr := config.GetEnvOrDefault("ADDRESS", *addrFlag)

	routerConfigurator := router.NewRouterConfigurator()
	r := routerConfigurator.SetupRoutes()

	srv := server.NewServer(r)
	if err := srv.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
