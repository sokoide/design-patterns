package usecase_test

import (
	"testing"

	"proxy-example/usecase"
)

type MockServer struct {
	CallCount int
}

func (m *MockServer) HandleRequest(url, method string) (int, string) {
	m.CallCount++
	return 200, "OK"
}

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(msg string) {
	m.Logs = append(m.Logs, msg)
}

func TestClient_SendRequests(t *testing.T) {
	logger := &MockLogger{}
	server := &MockServer{}
	client := usecase.NewClient(logger)

	client.SendRequests(server, "/test", 3)

	if server.CallCount != 3 {
		t.Errorf("expected 3 calls, got %d", server.CallCount)
	}

	if len(logger.Logs) != 3 {
		t.Errorf("expected 3 logs, got %d", len(logger.Logs))
	}

	expectedLog := "Request 1: Code=200, Body=OK"
	if logger.Logs[0] != expectedLog {
		t.Errorf("expected log '%s', got '%s'", expectedLog, logger.Logs[0])
	}
}
