package main

import (
	"net/http"
	"strconv"
	"strings"
)

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

var storage = MemStorage{make(map[string]float64), make(map[string]int64)}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", updateMetricsHandler)
	return http.ListenAndServe(`:8080`, mux)
}

func updateMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricType := parts[2]
	metricName := parts[3]
	metricValue := parts[4]

	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if metricType != "gauge" && metricType != "counter" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if metricType == "gauge" {
		metricValueFloat, _err := strconv.ParseFloat(metricValue, 64)
		if _err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		storage.gauge[metricName] = metricValueFloat
	}
	if metricType == "counter" {
		metricValueInt, _err := strconv.ParseInt(metricValue, 10, 64)
		if _err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		storage.counter[metricName] += metricValueInt
	}
	w.WriteHeader(http.StatusOK)
}
