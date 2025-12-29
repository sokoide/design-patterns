package usecase_test

import (
	"testing"

	"memento-example/domain"
	"memento-example/usecase"
)

type MockEditor struct {
	Content string
}

func (m *MockEditor) Type(words string) {
	m.Content += words
}

func (m *MockEditor) GetContent() string {
	return m.Content
}

func (m *MockEditor) CreateMemento() *domain.Memento {
	return &domain.Memento{State: m.Content}
}

func (m *MockEditor) Restore(mem *domain.Memento) {
	if mem != nil {
		m.Content = mem.State
	}
}

type MockLogger struct{}

func (m *MockLogger) Log(msg string) {}

func TestWriterService_WriteAndUndo(t *testing.T) {
	editor := &MockEditor{}
	logger := &MockLogger{}
	service := usecase.NewWriterService(editor, logger)

	service.Write("Hello")
	service.Save()
	service.Write(" World")

	if editor.GetContent() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", editor.GetContent())
	}

	service.Undo()

	if editor.GetContent() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", editor.GetContent())
	}
}
