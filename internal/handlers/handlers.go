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

//var tmpl = template.Must(template.ParseFiles("../../templates/metrics.html"))

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

	var valueStr string
	var err error

	switch metricType {
	case "gauge":
		value, ok := repo.GetGauge(metricName)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, err := fmt.Fprintf(w, "%f", value)
		if err != nil {
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
			return
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(valueStr))
	if err != nil {
		fmt.Println(err)
	}
}

func ListMetricsHandler(w http.ResponseWriter, _ *http.Request) {
	//gauges, counters := repo.GetAllMetrics()

	w.Header().Set("Content-Type", "text/html")

	//data := MetricsData{
	//	Gauges:   gauges,
	//	Counters: counters,
	//}
	//err := tmpl.Execute(w, data)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
}

func SetRepository(storage storage.Repository) {
	repo = storage
}
