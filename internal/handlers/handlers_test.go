package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepository struct {
	SaveGaugeCalled   bool
	SaveCounterCalled bool
}

func (m *MockRepository) SaveGauge(name string, value float64) {
	m.SaveGaugeCalled = true
}

func (m *MockRepository) SaveCounter(name string, value int64) {
	m.SaveCounterCalled = true
}

func TestUpdateMetricsHandler(t *testing.T) {
	mockRepo := &MockRepository{}
	SetRepository(mockRepo)

	testCases := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		gaugeCalled    bool
		counterCalled  bool
	}{
		{"ValidGauge", "POST", "/update/gauge/metricName/123.456", http.StatusOK, true, false},
		{"ValidCounter", "POST", "/update/counter/metricName/123", http.StatusOK, false, true},
		{"InvalidMethod", "GET", "/update/gauge/metricName/123.456", http.StatusMethodNotAllowed, false, false},
		{"InvalidMetricType", "POST", "/update/invalid/metricName/123", http.StatusBadRequest, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.url, nil)
			responseRecorder := httptest.NewRecorder()

			UpdateMetricsHandler(responseRecorder, request)

			if responseRecorder.Code != tc.expectedStatus {
				t.Errorf("Expected status %v, got %v", tc.expectedStatus, responseRecorder.Code)
			}

			if mockRepo.SaveGaugeCalled != tc.gaugeCalled {
				t.Errorf("Expected SaveGauge to be called: %v, got: %v", tc.gaugeCalled, mockRepo.SaveGaugeCalled)
			}

			if mockRepo.SaveCounterCalled != tc.counterCalled {
				t.Errorf("Expected SaveCounter to be called: %v, got: %v", tc.counterCalled, mockRepo.SaveCounterCalled)
			}

			mockRepo.SaveGaugeCalled = false
			mockRepo.SaveCounterCalled = false
		})
	}
}
