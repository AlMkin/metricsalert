package agent

import (
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"time"
)

type Agent struct {
	sender         sender.MetricSender
	collector      metrics.MetricCollector
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgent(s sender.MetricSender, c metrics.MetricCollector, poll time.Duration, report time.Duration) *Agent {
	return &Agent{
		sender:         s,
		collector:      c,
		pollInterval:   poll,
		reportInterval: report,
	}
}

func (a *Agent) Run() {
	pollTicker := time.NewTicker(a.pollInterval)
	reportTicker := time.NewTicker(a.reportInterval)
	defer pollTicker.Stop()
	defer reportTicker.Stop()

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
