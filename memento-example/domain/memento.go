package domain

type Memento struct {
	State string
}

func (m *Memento) GetSavedState() string {
	return m.State
}

type Editor interface {
	Type(words string)
	GetContent() string
	CreateMemento() *Memento
	Restore(m *Memento)
}

type Logger interface {
	Log(message string)
}
