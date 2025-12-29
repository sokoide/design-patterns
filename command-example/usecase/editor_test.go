package usecase_test

import (
	"strings"
	"testing"

	"command-example/domain"
	"command-example/usecase"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

type MockCommand struct {
	Executed bool
	Undone   bool
}

func (m *MockCommand) Do(b *domain.Buffer) {
	m.Executed = true
	b.Content += "X"
}

func (m *MockCommand) Undo(b *domain.Buffer) {
	m.Undone = true
	if len(b.Content) > 0 {
		b.Content = b.Content[:len(b.Content)-1]
	}
}

func TestEditor_Undo(t *testing.T) {
	logger := &MockLogger{}
	editor := usecase.NewEditor(logger)
	cmd := &MockCommand{}

	editor.Execute(cmd)
	if !cmd.Executed {
		t.Error("Command not executed")
	}
	if editor.GetContent() != "X" {
		t.Errorf("Buffer content wrong: %s", editor.GetContent())
	}

	editor.Undo()
	if !cmd.Undone {
		t.Error("Command not undone")
	}
	if editor.GetContent() != "" {
		t.Errorf("Buffer content not restored: %s", editor.GetContent())
	}

	// Undo on empty stack
	editor.Undo()
	if len(logger.Logs) == 0 || !strings.Contains(logger.Logs[0], "Nothing to undo") {
		t.Error("Expected warning log on empty undo")
	}
}
