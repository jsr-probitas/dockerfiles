package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	codeStr := chi.URLParam(r, "code")
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 100 || code > 599 {
		http.Error(w, "Invalid status code", http.StatusBadRequest)
		return
	}

	w.WriteHeader(code)
}
