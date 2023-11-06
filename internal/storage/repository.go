package storage

type Repository interface {
	SaveGauge(name string, value float64)
	SaveCounter(name string, value int64)
}
