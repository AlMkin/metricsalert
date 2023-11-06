package storage

type Repository interface {
	SaveGauge(name string, value float64)
	SaveCounter(name string, value int64)
	GetCounter(name string) (int64, bool)
	GetGauge(name string) (float64, bool)
	GetAllMetrics() (map[string]float64, map[string]int64)
}
