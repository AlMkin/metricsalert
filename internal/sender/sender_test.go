package sender_test

import (
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSenderSendOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	s := sender.NewSender(ts.URL)

	metricList := []metrics.Metric{
		{Type: "gauge", Name: "CPU", Value: 42.42},
	}

	s.Send(metricList)
}
