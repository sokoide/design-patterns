package adapter_test

import (
	"testing"

	"proxy-example/adapter"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

type MockServer struct {
	CalledCount int
}

func (m *MockServer) HandleRequest(url, method string) (int, string) {
	m.CalledCount++
	return 200, "OK"
}

func TestNginx_RateLimiting(t *testing.T) {
	logger := &MockLogger{}
	mockApp := &MockServer{}
	// Nginx default limit is 2
	nginx := adapter.NewNginx(mockApp, logger)

	// Request 1: Allowed
	code, _ := nginx.HandleRequest("/test", "GET")
	if code != 200 {
		t.Errorf("Request 1 blocked: %d", code)
	}

	// Request 2: Allowed
	code, _ = nginx.HandleRequest("/test", "GET")
	if code != 200 {
		t.Errorf("Request 2 blocked: %d", code)
	}

	// Request 3: Blocked
	code, _ = nginx.HandleRequest("/test", "GET")
	if code != 403 {
		t.Errorf("Request 3 allowed: %d", code)
	}

	// Verify Application was only called twice
	if mockApp.CalledCount != 2 {
		t.Errorf("AppServer called %d times, expected 2", mockApp.CalledCount)
	}
}
