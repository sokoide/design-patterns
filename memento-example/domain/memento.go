package domain

// Memento holds the saved state for the editor.
type Memento struct {
	state string
}

// NewMemento creates a memento for a given state.
func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

// State returns the saved state.
func (m *Memento) State() string {
	return m.state
}

// Editor defines the originator behavior.
type Editor interface {
	Type(words string)
	GetContent() string
	CreateMemento() *Memento
	Restore(m *Memento)
}

// Logger abstracts logging for the domain.
type Logger interface {
	Log(message string)
}
