package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestStatusHandler(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		expectedStatus int
	}{
		{
			name:           "200 returns 200 OK",
			code:           "200",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "404 returns 404",
			code:           "404",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "500 returns 500",
			code:           "500",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "99 returns 400 (below valid range)",
			code:           "99",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "600 returns 400 (above valid range)",
			code:           "600",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "abc returns 400 (non-numeric)",
			code:           "abc",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/status/{code}", StatusHandler)

			req := httptest.NewRequest(http.MethodGet, "/status/"+tt.code, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
