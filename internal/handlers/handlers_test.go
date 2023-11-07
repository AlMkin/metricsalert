package handlers

import (
	"github.com/AlMkin/metricsalert/internal/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func pointerToInt64(value int64) *int64 {
	return &value
}

func pointerToFloat64(value float64) *float64 {
	return &value
}

func TestUpdateMetricsHandler(t *testing.T) {
	memStorage := storage.NewMemStorage()
	SetRepository(memStorage)

	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", UpdateMetricsHandler)

	testCases := []struct {
		name            string
		method          string
		url             string
		expectedStatus  int
		expectedGauge   *float64
		expectedCounter *int64
	}{
		{"ValidGauge", "POST", "/update/gauge/metricName/123.456", http.StatusOK, pointerToFloat64(123.456), nil},
		{"ValidCounter", "POST", "/update/counter/metricName/123", http.StatusOK, nil, pointerToInt64(123)},
		{"InvalidMethod", "GET", "/update/gauge/metricName/123.456", http.StatusMethodNotAllowed, nil, nil},
		{"InvalidMetricType", "POST", "/update/invalid/metricName/123", http.StatusBadRequest, nil, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.url, nil)
			responseRecorder := httptest.NewRecorder()

			r.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.expectedStatus {
				t.Errorf("Expected status %v, got %v", tc.expectedStatus, responseRecorder.Code)
			}

			if tc.expectedGauge != nil {
				value, ok := memStorage.GetGauge("metricName")
				if !ok || *tc.expectedGauge != value {
					t.Errorf("Expected gauge value to be %v, got %v", *tc.expectedGauge, value)
				}
			}

			if tc.expectedCounter != nil {
				value, ok := memStorage.GetCounter("metricName")
				if !ok || *tc.expectedCounter != value {
					t.Errorf("Expected counter value to be %v, got %v", *tc.expectedCounter, value)
				}
			}

			// Сброс состояния хранилища перед следующим тестом
			memStorage = storage.NewMemStorage()
			SetRepository(memStorage)
		})
	}
}

func TestGetMetricsHandler(t *testing.T) {
	memStorage := storage.NewMemStorage()
	SetRepository(memStorage)

	r := chi.NewRouter()
	r.Get("/value/{type}/{name}", GetMetricsHandler)

	// Добавляем реальные значения в хранилище
	memStorage.SaveGauge("testGauge", 42.42)
	memStorage.SaveCounter("testCounter", 100)

	testCases := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{"ExistingGauge", "/value/gauge/testGauge", http.StatusOK, "42.420000"},
		{"NonExistingGauge", "/value/gauge/nonExisting", http.StatusNotFound, ""},
		{"ExistingCounter", "/value/counter/testCounter", http.StatusOK, "100"},
		{"NonExistingCounter", "/value/counter/nonExisting", http.StatusNotFound, ""},
		{"InvalidMetricType", "/value/invalid/testMetric", http.StatusBadRequest, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tc.url, nil)
			responseRecorder := httptest.NewRecorder()

			r.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.expectedBody {
				t.Errorf("Expected body to be %s, got %s", tc.expectedBody, responseRecorder.Body.String())
			}
		})
	}
}
