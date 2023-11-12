package storage

// Repository defines an interface for a metrics storage.
// This interface allows for saving and retrieving various types of metrics,
// such as counters and gauges.
type Repository interface {
	// SaveGauge saves a gauge type metric.
	// name - a unique name of the metric.
	// value - the value of the metric.
	SaveGauge(name string, value float64)

	// SaveCounter increments the value of a counter type metric.
	// name - a unique name of the counter.
	// value - the amount by which the counter is increased.
	SaveCounter(name string, value int64)

	// GetCounter returns the current value of a counter and a flag indicating
	// whether the counter exists.
	// name - the name of the counter.
	// Returns the value of the counter and a boolean indicating its existence.
	GetCounter(name string) (int64, bool)

	// GetGauge returns the current value of a gauge type metric and a flag indicating
	// whether the gauge exists.
	// name - the name of the gauge metric.
	// Returns the value of the metric and a boolean indicating its existence.
	GetGauge(name string) (float64, bool)

	// GetAllMetrics returns all stored metrics.
	// Returns two maps: one for Gauge values and one for Counter values.
	GetAllMetrics() (map[string]float64, map[string]int64)
}
