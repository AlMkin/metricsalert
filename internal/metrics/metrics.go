package metrics

import (
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"math/rand"
)

type MetricCollector interface {
	Collect()
	GetMetrics() []Metric
	ResetMetrics()
}

type Collector struct {
	metrics       []Metric
	pollCount     int64
	randomValue   float64
	metricsGetter Getter // Добавляем это поле для инъекции зависимостей
}

type Getter interface {
	GetRuntimeMetrics() runtimeinfo.Stats
}

func NewCollector(metricsGetter Getter) *Collector {
	return &Collector{
		randomValue:   rand.Float64(), // Инициализация случайным значением
		metricsGetter: metricsGetter,  // Используем инъекцию зависимостей
	}
}

type Metric struct {
	Type  string
	Name  string
	Value float64
}

var _ MetricCollector = (*Collector)(nil)

func (c *Collector) Collect() {
	stats := c.metricsGetter.GetRuntimeMetrics()
	c.pollCount++
	c.randomValue = rand.Float64()
	c.metrics = append(c.metrics, convertToMetricSlice(stats, c.pollCount, c.randomValue)...)
}

func (c *Collector) GetMetrics() []Metric {
	return c.metrics
}

func (c *Collector) ResetMetrics() {
	c.metrics = nil
}

func convertToMetricSlice(memStats runtimeinfo.Stats, pollCount int64, randomValue float64) []Metric {
	metrics := []Metric{
		{Type: "gauge", Name: "Alloc", Value: float64(memStats.Alloc)},
		{Type: "gauge", Name: "BuckHashSys", Value: float64(memStats.BuckHashSys)},
		{Type: "gauge", Name: "Frees", Value: float64(memStats.Frees)},
		{Type: "gauge", Name: "GCCPUFraction", Value: float64(memStats.GCCPUFraction)},
		{Type: "gauge", Name: "GCSys", Value: float64(memStats.GCSys)},
		{Type: "gauge", Name: "HeapAlloc", Value: float64(memStats.HeapAlloc)},
		{Type: "gauge", Name: "HeapIdle", Value: float64(memStats.HeapIdle)},
		{Type: "gauge", Name: "HeapInuse", Value: float64(memStats.HeapInuse)},
		{Type: "gauge", Name: "HeapObjects", Value: float64(memStats.HeapObjects)},
		{Type: "gauge", Name: "HeapReleased", Value: float64(memStats.HeapReleased)},
		{Type: "gauge", Name: "HeapSys", Value: float64(memStats.HeapSys)},
		{Type: "gauge", Name: "LastGC", Value: float64(memStats.LastGC)},
		{Type: "gauge", Name: "Lookups", Value: float64(memStats.Lookups)},
		{Type: "gauge", Name: "MCacheInuse", Value: float64(memStats.MCacheInuse)},
		{Type: "gauge", Name: "MCacheSys", Value: float64(memStats.MCacheSys)},
		{Type: "gauge", Name: "MSpanInuse", Value: float64(memStats.MSpanInuse)},
		{Type: "gauge", Name: "MSpanSys", Value: float64(memStats.MSpanSys)},
		{Type: "gauge", Name: "Mallocs", Value: float64(memStats.Mallocs)},
		{Type: "gauge", Name: "NextGC", Value: float64(memStats.NextGC)},
		{Type: "gauge", Name: "NumForcedGC", Value: float64(memStats.NumForcedGC)},
		{Type: "gauge", Name: "NumGC", Value: float64(memStats.NumGC)},
		{Type: "gauge", Name: "OtherSys", Value: float64(memStats.OtherSys)},
		{Type: "gauge", Name: "PauseTotalNs", Value: float64(memStats.PauseTotalNs)},
		{Type: "gauge", Name: "StackInuse", Value: float64(memStats.StackInuse)},
		{Type: "gauge", Name: "StackSys", Value: float64(memStats.StackSys)},
		{Type: "gauge", Name: "Sys", Value: float64(memStats.Sys)},
		{Type: "gauge", Name: "TotalAlloc", Value: float64(memStats.TotalAlloc)},
	}
	metrics = append(metrics, Metric{Type: "counter", Name: "PollCount", Value: float64(pollCount)})
	metrics = append(metrics, Metric{Type: "gauge", Name: "RandomValue", Value: randomValue})
	return metrics
}
