package storage

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{make(map[string]float64), make(map[string]int64)}
}

func (s *MemStorage) SaveGauge(name string, value float64) {
	s.gauge[name] = value
}

func (s *MemStorage) SaveCounter(name string, value int64) {
	s.counter[name] += value
}
