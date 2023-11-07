package handlers

import (
	"fmt"
	"github.com/AlMkin/metricsalert/internal/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

var repo storage.Repository

type MetricsData struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

func UpdateMetricsHandler(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "name")
	metricValue := chi.URLParam(r, "value")

	switch metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		repo.SaveGauge(metricName, value)

	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		repo.SaveCounter(metricName, value)

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "name")

	switch metricType {
	case "gauge":
		value, ok := repo.GetGauge(metricName)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, err := fmt.Fprintf(w, "%g", value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "counter":
		value, ok := repo.GetCounter(metricName)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, err := fmt.Fprintf(w, "%d", value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func ListMetricsHandler(w http.ResponseWriter, _ *http.Request) {
	gauges, counters := repo.GetAllMetrics()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	htmlStart := `<!DOCTYPE html><html lang="en"><body><h1>Metrics</h1><ul>`
	htmlGauges := ""
	for name, value := range gauges {
		htmlGauges += fmt.Sprintf("<li>%s (gauge): %.2f</li>", name, value)
	}
	htmlCounters := ""
	for name, value := range counters {
		htmlCounters += fmt.Sprintf("<li>%s (counter): %d</li>", name, value)
	}
	htmlEnd := `</ul></body></html>`

	fullHTML := htmlStart + htmlGauges + htmlCounters + htmlEnd // Corrected variable name

	_, err := w.Write([]byte(fullHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SetRepository(storage storage.Repository) {
	repo = storage
}
