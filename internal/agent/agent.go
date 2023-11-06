package agent

import (
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

type Agent struct {
	sender    sender.MetricSender
	collector metrics.MetricCollector
}

func NewAgent(s sender.MetricSender, c metrics.MetricCollector) *Agent {
	return &Agent{
		sender:    s,
		collector: c,
	}
}

func (a *Agent) Run() {
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	for {
		select {
		case <-pollTicker.C:
			a.collector.Collect()
		case <-reportTicker.C:
			getMetrics := a.collector.GetMetrics()
			a.sender.Send(getMetrics)
			a.collector.ResetMetrics()
		}
	}
}
