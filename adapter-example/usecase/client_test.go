package usecase_test

import (
	"testing"

	"adapter-example/usecase"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

type MockComputer struct {
	Inserted bool
}

func (m *MockComputer) InsertIntoLightningPort() {
	m.Inserted = true
}

func TestClient_InsertLightningConnectorIntoComputer(t *testing.T) {
	logger := &MockLogger{}
	client := usecase.NewClient(logger)
	computer := &MockComputer{}

	client.InsertLightningConnectorIntoComputer(computer)

	if !computer.Inserted {
		t.Error("expected computer to be inserted")
	}

	if len(logger.Logs) != 1 {
		t.Errorf("expected 1 log, got %d", len(logger.Logs))
	}
}
