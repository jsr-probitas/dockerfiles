package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeadersHandler(t *testing.T) {
	t.Run("returns request headers correctly", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/headers", nil)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en-US")
		rec := httptest.NewRecorder()

		HeadersHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		contentType := rec.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", contentType)
		}

		var resp HeadersResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Headers["Accept"] != "application/json" {
			t.Errorf("expected Accept=application/json, got %s", resp.Headers["Accept"])
		}

		if resp.Headers["Accept-Language"] != "en-US" {
			t.Errorf("expected Accept-Language=en-US, got %s", resp.Headers["Accept-Language"])
		}
	})

	t.Run("custom headers are echoed back", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/headers", nil)
		req.Header.Set("X-Custom-Header", "custom-value")
		req.Header.Set("X-Request-Id", "12345")
		req.Header.Set("Authorization", "Bearer token123")
		rec := httptest.NewRecorder()

		HeadersHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp HeadersResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Headers["X-Custom-Header"] != "custom-value" {
			t.Errorf("expected X-Custom-Header=custom-value, got %s", resp.Headers["X-Custom-Header"])
		}

		if resp.Headers["X-Request-Id"] != "12345" {
			t.Errorf("expected X-Request-Id=12345, got %s", resp.Headers["X-Request-Id"])
		}

		if resp.Headers["Authorization"] != "Bearer token123" {
			t.Errorf("expected Authorization=Bearer token123, got %s", resp.Headers["Authorization"])
		}
	})
}
