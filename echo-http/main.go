package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jsr-probitas/dockerfiles/echo-http/handlers"
)

func main() {
	cfg := LoadConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Echo endpoints
	r.Get("/get", handlers.EchoHandler)
	r.Post("/post", handlers.EchoHandler)
	r.Put("/put", handlers.EchoHandler)
	r.Patch("/patch", handlers.EchoHandler)
	r.Delete("/delete", handlers.EchoHandler)

	// Utility endpoints
	r.Get("/headers", handlers.HeadersHandler)
	r.Get("/status/{code}", handlers.StatusHandler)
	r.Get("/delay/{seconds}", handlers.DelayHandler)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Starting server on %s", cfg.Addr())
	if err := http.ListenAndServe(cfg.Addr(), r); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
