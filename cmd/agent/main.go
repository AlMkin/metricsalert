package main

import (
	"github.com/AlMkin/metricsalert/internal/agent"
	"github.com/AlMkin/metricsalert/internal/config"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"log"
)

func main() {
	cfg := config.GetConfig()

	newSender := sender.NewSender(cfg.Address)
	metricsGetter := &runtimeinfo.Getter{}

	collector := metrics.NewCollector(metricsGetter)
	a := agent.NewAgent(newSender, collector, cfg.PollInterval, cfg.ReportInterval)

	log.Println("Agent started with the following parameters:")
	log.Printf("ADDRESS: %s", cfg.Address)
	log.Printf("REPORT_INTERVAL: %v", cfg.ReportInterval)
	log.Printf("POLL_INTERVAL: %v", cfg.PollInterval)

	a.Run()
}
