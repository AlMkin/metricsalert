package main

import (
	"github.com/AlMkin/metricsalert/internal/agent"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
)

func main() {
	newSender := sender.NewSender("http://localhost:8080")
	metricsGetter := &runtimeinfo.Getter{}

	collector := metrics.NewCollector(metricsGetter)
	a := agent.NewAgent(newSender, collector)
	a.Run()
}
