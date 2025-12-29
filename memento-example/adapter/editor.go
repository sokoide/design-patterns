package adapter

import (
	"memento-example/domain"
)

// Originator
type Editor struct {
	content string
}

func (e *Editor) Type(words string) {
	e.content = e.content + " " + words
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) CreateMemento() *domain.Memento {
	return &domain.Memento{State: e.content}
}

func (e *Editor) Restore(m *domain.Memento) {
	if m == nil {
		return
	}
	e.content = m.GetSavedState()
}
