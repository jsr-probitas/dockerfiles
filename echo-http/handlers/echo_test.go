package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandler_GET(t *testing.T) {
	t.Run("returns correct method and url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get?foo=bar&baz=qux", nil)
		req.Header.Set("X-Custom-Header", "test-value")
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Method != "GET" {
			t.Errorf("expected method GET, got %s", resp.Method)
		}

		if resp.URL != "/get?foo=bar&baz=qux" {
			t.Errorf("expected url /get?foo=bar&baz=qux, got %s", resp.URL)
		}

		if resp.Args["foo"] != "bar" {
			t.Errorf("expected args[foo]=bar, got %s", resp.Args["foo"])
		}

		if resp.Args["baz"] != "qux" {
			t.Errorf("expected args[baz]=qux, got %s", resp.Args["baz"])
		}

		if resp.Headers["X-Custom-Header"] != "test-value" {
			t.Errorf("expected header X-Custom-Header=test-value, got %s", resp.Headers["X-Custom-Header"])
		}
	})
}

func TestEchoHandler_POST(t *testing.T) {
	t.Run("with JSON body returns correct json field", func(t *testing.T) {
		body := `{"message":"hello","count":42}`
		req := httptest.NewRequest(http.MethodPost, "/post", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Method != "POST" {
			t.Errorf("expected method POST, got %s", resp.Method)
		}

		if resp.Data != body {
			t.Errorf("expected data %s, got %s", body, resp.Data)
		}

		jsonMap, ok := resp.JSON.(map[string]any)
		if !ok {
			t.Fatalf("expected JSON to be a map, got %T", resp.JSON)
		}

		if jsonMap["message"] != "hello" {
			t.Errorf("expected json.message=hello, got %v", jsonMap["message"])
		}

		if jsonMap["count"] != float64(42) {
			t.Errorf("expected json.count=42, got %v", jsonMap["count"])
		}
	})

	t.Run("with form data returns correct form field", func(t *testing.T) {
		body := "name=john&age=30"
		req := httptest.NewRequest(http.MethodPost, "/post", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Form["name"] != "john" {
			t.Errorf("expected form[name]=john, got %s", resp.Form["name"])
		}

		if resp.Form["age"] != "30" {
			t.Errorf("expected form[age]=30, got %s", resp.Form["age"])
		}
	})
}

func TestEchoHandler_PUT(t *testing.T) {
	t.Run("works correctly", func(t *testing.T) {
		body := `{"update":"data"}`
		req := httptest.NewRequest(http.MethodPut, "/put", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Method != "PUT" {
			t.Errorf("expected method PUT, got %s", resp.Method)
		}

		if resp.Data != body {
			t.Errorf("expected data %s, got %s", body, resp.Data)
		}
	})
}

func TestEchoHandler_PATCH(t *testing.T) {
	t.Run("works correctly", func(t *testing.T) {
		body := `{"patch":"value"}`
		req := httptest.NewRequest(http.MethodPatch, "/patch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Method != "PATCH" {
			t.Errorf("expected method PATCH, got %s", resp.Method)
		}
	})
}

func TestEchoHandler_DELETE(t *testing.T) {
	t.Run("works correctly", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/delete", nil)
		rec := httptest.NewRecorder()

		EchoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var resp EchoResponse
		if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.Method != "DELETE" {
			t.Errorf("expected method DELETE, got %s", resp.Method)
		}
	})
}
