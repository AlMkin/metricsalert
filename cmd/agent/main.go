package main

import (
	"github.com/AlMkin/metricsalert/internal/agent"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"github.com/AlMkin/metricsalert/pkg/config"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Agent failed to load config: %v", err)
	}

	newSender := sender.NewSender(cfg.Address)
	metricsGetter := &runtimeinfo.Getter{}
	collector := metrics.NewCollector(metricsGetter)

	a := agent.NewAgent(newSender, collector, cfg.PollInterval, cfg.ReportInterval)

	a.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	a.Stop()
}
