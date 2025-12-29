package adapter_test

import (
	"strings"
	"testing"

	"observer-example/adapter"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestEmailNotifier_OnUpdate(t *testing.T) {
	logger := &MockLogger{}
	notifier := adapter.NewEmailNotifier("test@example.com", logger)
	notifier.OnUpdate("Something happened")

	if len(logger.Logs) != 1 {
		t.Errorf("expected 1 log, got %d", len(logger.Logs))
	}
	if !strings.Contains(logger.Logs[0], "Email to test@example.com") {
		t.Error("log missing email prefix")
	}
}
