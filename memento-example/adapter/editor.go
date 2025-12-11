package adapter

import (
	"fmt"
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
		fmt.Println("[WARN] no memento to restore")
		return
	}
	e.content = m.GetSavedState()
}

// Caretaker
type Caretaker struct {
	mementoArray []*domain.Memento
}

func (c *Caretaker) AddMemento(m *domain.Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) GetMemento(index int) *domain.Memento {
	if index < 0 || index >= len(c.mementoArray) {
		return nil
	}
	return c.mementoArray[index]
}
