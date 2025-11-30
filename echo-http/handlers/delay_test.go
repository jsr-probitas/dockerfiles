package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

func TestDelayHandler(t *testing.T) {
	t.Run("delay/0 returns immediately", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/delay/{seconds}", DelayHandler)

		req := httptest.NewRequest(http.MethodGet, "/delay/0", nil)
		rec := httptest.NewRecorder()

		start := time.Now()
		r.ServeHTTP(rec, req)
		elapsed := time.Since(start)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		if elapsed > 100*time.Millisecond {
			t.Errorf("expected immediate response, took %v", elapsed)
		}

		var resp map[string]any
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp["delay"] != float64(0) {
			t.Errorf("expected delay=0, got %v", resp["delay"])
		}
	})

	t.Run("delay/1 takes approximately 1 second", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/delay/{seconds}", DelayHandler)

		req := httptest.NewRequest(http.MethodGet, "/delay/1", nil)
		rec := httptest.NewRecorder()

		start := time.Now()
		r.ServeHTTP(rec, req)
		elapsed := time.Since(start)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		if elapsed < 900*time.Millisecond || elapsed > 1200*time.Millisecond {
			t.Errorf("expected delay of ~1 second, took %v", elapsed)
		}

		var resp map[string]any
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp["delay"] != float64(1) {
			t.Errorf("expected delay=1, got %v", resp["delay"])
		}
	})

	t.Run("delay/-1 returns 400 (negative value)", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/delay/{seconds}", DelayHandler)

		req := httptest.NewRequest(http.MethodGet, "/delay/-1", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", rec.Code)
		}
	})

	t.Run("delay/abc returns 400 (non-numeric)", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/delay/{seconds}", DelayHandler)

		req := httptest.NewRequest(http.MethodGet, "/delay/abc", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", rec.Code)
		}
	})
}
