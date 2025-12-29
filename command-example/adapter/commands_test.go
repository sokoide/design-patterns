package adapter_test

import (
	"testing"

	"command-example/adapter"
	"command-example/domain"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestInsertCommand(t *testing.T) {
	logger := &MockLogger{}
	buffer := domain.NewBuffer()
	cmd := adapter.NewInsertCommand("Hello", logger)

	cmd.Do(buffer)
	if buffer.Content != "Hello" {
		t.Errorf("Insert failed: %s", buffer.Content)
	}

	cmd.Undo(buffer)
	if buffer.Content != "" {
		t.Errorf("Undo Insert failed: %s", buffer.Content)
	}
}

func TestDeleteCommand(t *testing.T) {
	logger := &MockLogger{}
	buffer := domain.NewBuffer()
	buffer.Content = "HelloWorld"

	// Delete "World" (5 chars)
	cmd := adapter.NewDeleteCommand(5, logger)

	cmd.Do(buffer)
	if buffer.Content != "Hello" {
		t.Errorf("Delete failed: %s", buffer.Content)
	}

	cmd.Undo(buffer)
	if buffer.Content != "HelloWorld" {
		t.Errorf("Undo Delete failed: %s", buffer.Content)
	}
}
