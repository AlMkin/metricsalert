package handlers

import (
	"github.com/AlMkin/metricsalert/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

var repo storage.Repository

func UpdateMetricsHandler(w http.ResponseWriter, r *http.Request) {
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
		repo.SaveGauge(metricName, metricValueFloat)
	}
	if metricType == "counter" {
		metricValueInt, _err := strconv.ParseInt(metricValue, 10, 64)
		if _err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		repo.SaveCounter(metricName, metricValueInt)
	}
	w.WriteHeader(http.StatusOK)
}

func SetRepository(storage storage.Repository) {
	repo = storage
}
