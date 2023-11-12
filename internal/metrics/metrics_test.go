package metrics_test

import (
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"testing"
	"time"
)

type MockGetter struct {
	mock.Mock
}

func (m *MockGetter) GetRuntimeMetrics() runtimeinfo.Stats {
	args := m.Called()
	return args.Get(0).(runtimeinfo.Stats)
}

// TestCollectorCollect проверяет, что метод Collect правильно собирает метрики.
func TestCollectorCollect(t *testing.T) {
	mockGetter := new(MockGetter)
	rand.NewSource(time.Now().UnixNano())
	expectedStats := runtimeinfo.Stats{
		Alloc:      1024,
		TotalAlloc: 2048,
	}
	mockGetter.On("GetRuntimeMetrics").Return(expectedStats)

	collector := metrics.NewCollector(mockGetter)
	collector.Collect()

	getMetrics := collector.GetMetrics()
	assert.NotEmpty(t, getMetrics, "Metrics should not be empty after collection")

	// Проверяем, что метрика PollCount увеличивается с каждым вызовом Collect.
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
