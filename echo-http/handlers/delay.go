package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

const maxDelaySeconds = 30

func DelayHandler(w http.ResponseWriter, r *http.Request) {
	secondsStr := chi.URLParam(r, "seconds")
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil || seconds < 0 {
		http.Error(w, "Invalid delay value", http.StatusBadRequest)
		return
	}

	if seconds > maxDelaySeconds {
		seconds = maxDelaySeconds
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	response := map[string]any{
		"delay":  seconds,
		"method": r.Method,
		"url":    r.URL.RequestURI(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
