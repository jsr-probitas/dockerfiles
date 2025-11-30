package server

import (
	"testing"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func TestNewHealthServer_SetsInitialServingStatus(t *testing.T) {
	h := NewHealthServer()

	// Overall server status should be SERVING
	if status := h.GetServingStatus(""); status != healthpb.HealthCheckResponse_SERVING {
		t.Errorf("expected overall status SERVING, got %v", status)
	}

	// Echo service status should be SERVING
	if status := h.GetServingStatus("echo.v1.Echo"); status != healthpb.HealthCheckResponse_SERVING {
		t.Errorf("expected echo.v1.Echo status SERVING, got %v", status)
	}
}

func TestHealthServer_SetServingStatus(t *testing.T) {
	h := NewHealthServer()

	h.SetServingStatus("test.service", healthpb.HealthCheckResponse_SERVING)
	if status := h.GetServingStatus("test.service"); status != healthpb.HealthCheckResponse_SERVING {
		t.Errorf("expected SERVING, got %v", status)
	}

	h.SetServingStatus("test.service", healthpb.HealthCheckResponse_NOT_SERVING)
	if status := h.GetServingStatus("test.service"); status != healthpb.HealthCheckResponse_NOT_SERVING {
		t.Errorf("expected NOT_SERVING, got %v", status)
	}
}

func TestHealthServer_GetServingStatus_UnknownService(t *testing.T) {
	h := NewHealthServer()

	status := h.GetServingStatus("unknown.service")
	if status != healthpb.HealthCheckResponse_SERVICE_UNKNOWN {
		t.Errorf("expected SERVICE_UNKNOWN for unregistered service, got %v", status)
	}
}

func TestHealthServer_Shutdown(t *testing.T) {
	h := NewHealthServer()

	// Add additional service
	h.SetServingStatus("test.service", healthpb.HealthCheckResponse_SERVING)

	// Shutdown
	h.Shutdown()

	// All services should be NOT_SERVING after shutdown
	if status := h.GetServingStatus(""); status != healthpb.HealthCheckResponse_NOT_SERVING {
		t.Errorf("expected overall status NOT_SERVING after shutdown, got %v", status)
	}
	if status := h.GetServingStatus("echo.v1.Echo"); status != healthpb.HealthCheckResponse_NOT_SERVING {
		t.Errorf("expected echo.v1.Echo status NOT_SERVING after shutdown, got %v", status)
	}
	if status := h.GetServingStatus("test.service"); status != healthpb.HealthCheckResponse_NOT_SERVING {
		t.Errorf("expected test.service status NOT_SERVING after shutdown, got %v", status)
	}
}
