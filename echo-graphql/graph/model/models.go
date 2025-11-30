package model

import (
	"context"
	"net/http"
)

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

type EchoResult struct {
	Message *string `json:"message,omitempty"`
	Error   *string `json:"error,omitempty"`
}

// Headers represents HTTP request headers for echoHeaders query
type Headers struct {
	Request *http.Request `json:"-"`
}

// HeaderEntry represents a single header key-value pair
type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NestedEcho represents a nested echo structure for testing recursive parsing
type NestedEcho struct {
	Value string      `json:"value"`
	Child *NestedEcho `json:"child,omitempty"`
}

// EchoListItem represents a single item in an echo list
type EchoListItem struct {
	Index   int    `json:"index"`
	Message string `json:"message"`
}

// Key for storing http.Request in context
type contextKey string

const RequestKey contextKey = "httpRequest"

// GetRequestFromContext retrieves the http.Request from context
func GetRequestFromContext(ctx context.Context) *http.Request {
	if req, ok := ctx.Value(RequestKey).(*http.Request); ok {
		return req
	}
	return nil
}
