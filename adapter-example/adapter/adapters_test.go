package adapter_test

import (
	"testing"

	"adapter-example/adapter"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestMac_InsertIntoLightningPort(t *testing.T) {
	logger := &MockLogger{}
	mac := adapter.NewMac(logger)

	mac.InsertIntoLightningPort()

	if len(logger.Logs) != 1 {
		t.Errorf("expected 1 log, got %d", len(logger.Logs))
	}
	expected := "Lightning connector is plugged into Mac machine."
	if logger.Logs[0] != expected {
		t.Errorf("expected %q, got %q", expected, logger.Logs[0])
	}
}

func TestWindowsAdapter_InsertIntoLightningPort(t *testing.T) {
	logger := &MockLogger{}
	windows := adapter.NewWindows(logger)
	windowsAdapter := adapter.NewWindowsAdapter(windows, logger)

	windowsAdapter.InsertIntoLightningPort()

	if len(logger.Logs) != 2 {
		t.Errorf("expected 2 logs (adapter + windows), got %d", len(logger.Logs))
	}
	expectedAdapter := "Adapter converts Lightning signal to USB."
	if logger.Logs[0] != expectedAdapter {
		t.Errorf("expected %q, got %q", expectedAdapter, logger.Logs[0])
	}
	expectedWindows := "USB connector is plugged into Windows machine."
	if logger.Logs[1] != expectedWindows {
		t.Errorf("expected %q, got %q", expectedWindows, logger.Logs[1])
	}
}
