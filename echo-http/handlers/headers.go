package handlers

import (
	"encoding/json"
	"net/http"
)

type HeadersResponse struct {
	Headers map[string]string `json:"headers"`
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	response := HeadersResponse{
		Headers: make(map[string]string),
	}

	for key, values := range r.Header {
		if len(values) > 0 {
			response.Headers[key] = values[0]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
