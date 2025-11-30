package server

import (
	"sync"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// HealthServer wraps the standard gRPC health server with service management.
type HealthServer struct {
	*health.Server
	mu       sync.RWMutex
	services map[string]healthpb.HealthCheckResponse_ServingStatus
}

// NewHealthServer creates a new health server with default services.
func NewHealthServer() *HealthServer {
	h := &HealthServer{
		Server:   health.NewServer(),
		services: make(map[string]healthpb.HealthCheckResponse_ServingStatus),
	}

	// Set overall server status (empty service name = overall status)
	h.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// Set Echo service status
	h.SetServingStatus("echo.v1.Echo", healthpb.HealthCheckResponse_SERVING)

	return h
}

// SetServingStatus updates the serving status for a service.
func (h *HealthServer) SetServingStatus(service string, status healthpb.HealthCheckResponse_ServingStatus) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.services[service] = status
	h.Server.SetServingStatus(service, status)
}

// GetServingStatus returns the current serving status for a service.
func (h *HealthServer) GetServingStatus(service string) healthpb.HealthCheckResponse_ServingStatus {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if status, ok := h.services[service]; ok {
		return status
	}
	return healthpb.HealthCheckResponse_SERVICE_UNKNOWN
}

// Shutdown sets all services to NOT_SERVING status.
func (h *HealthServer) Shutdown() {
	h.mu.Lock()
	defer h.mu.Unlock()
	for service := range h.services {
		h.services[service] = healthpb.HealthCheckResponse_NOT_SERVING
		h.Server.SetServingStatus(service, healthpb.HealthCheckResponse_NOT_SERVING)
	}
}
