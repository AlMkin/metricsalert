package storage

// MemStorage теперь приватный
type memStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() Repository {
	return &memStorage{make(map[string]float64), make(map[string]int64)}
}

func (s *memStorage) SaveGauge(name string, value float64) {
	s.gauge[name] = value
}

func (s *memStorage) SaveCounter(name string, value int64) {
	s.counter[name] += value
}

func (s *memStorage) GetGauge(name string) (float64, bool) {
	value, ok := s.gauge[name]
	return value, ok
}

func (s *memStorage) GetCounter(name string) (int64, bool) {
	value, ok := s.counter[name]
	return value, ok
}

func (s *memStorage) GetAllMetrics() (map[string]float64, map[string]int64) {
	return s.gauge, s.counter
}
