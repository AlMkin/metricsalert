package sender

import (
	"bytes"
	"fmt"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"net/http"
)

type Sender struct {
	serverAddress string
}

type MetricSender interface {
	Send(metrics []metrics.Metric)
}

var _ MetricSender = (*Sender)(nil)

func NewSender(serverAddress string) *Sender {
	return &Sender{
		serverAddress: serverAddress,
	}
}

func (s *Sender) Send(metrics []metrics.Metric) {
	for _, m := range metrics {
		body := bytes.NewBufferString(fmt.Sprintf("%s/update/%s/%s/%f", s.serverAddress, m.Type, m.Name, m.Value))
		_, err := http.Post(body.String(), "text/plain", nil)
		if err != nil {
			fmt.Println("Error sending metrics:", err)
			continue
		}
	}
}
