package adapter

import "memento-example/domain"

// Editor is the originator that produces and consumes mementos.
type Editor struct {
	content string
}

func (e *Editor) Type(words string) {
	if e.content == "" {
		e.content = words
		return
	}
	e.content += " " + words
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) CreateMemento() *domain.Memento {
	return domain.NewMemento(e.content)
}

func (e *Editor) Restore(m *domain.Memento) {
	if m == nil {
		return
	}
	e.content = m.State()
}
