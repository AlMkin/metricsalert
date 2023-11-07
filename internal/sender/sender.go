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

func NewSender(serverAddress string) *Sender {
	return &Sender{
		serverAddress: serverAddress,
	}
}

var _ MetricSender = (*Sender)(nil)

func (s *Sender) Send(metrics []metrics.Metric) {
	for _, m := range metrics {
		body := bytes.NewBufferString(fmt.Sprintf("%s/update/%s/%s/%f", s.serverAddress, m.Type, m.Name, m.Value))
		response, err := http.Post(body.String(), "text/plain", nil)
		if err != nil {
			fmt.Println("Error sending metrics:", err)
			continue
		}
		err = response.Body.Close()
		if err != nil {
			return
		}
	}
}
