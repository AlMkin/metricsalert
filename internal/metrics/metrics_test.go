package metrics_test

import (
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/mocks"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"math/rand"
	"testing"
	"time"
)

func TestCollectorCollect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetter := mock_getter.NewMockGetter(ctrl)
	rand.NewSource(time.Now().UnixNano())
	expectedStats := runtimeinfo.Stats{
		Alloc:      1024,
		TotalAlloc: 2048,
	}
	mockGetter.EXPECT().GetRuntimeMetrics().Return(expectedStats).Times(1)

	collector := metrics.NewCollector(mockGetter)
	collector.Collect()

	getMetrics := collector.GetMetrics()
	assert.NotEmpty(t, getMetrics, "Metrics should not be empty after collection")

	expectedPollCount := 1.0
	foundPollCountMetric := false
	delta := 0.001
	for _, metric := range getMetrics {
		if metric.Name == "PollCount" {
			foundPollCountMetric = true
			assert.InDelta(t, expectedPollCount, metric.Value, delta, "PollCount metric value should be within delta")
			break
		}
	}
	assert.True(t, foundPollCountMetric, "Should find a PollCount metric")
}
