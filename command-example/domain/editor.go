package domain

// Buffer is the Receiver in the Command Pattern.
// It holds the actual state (the text content).
type Buffer struct {
	Content string
}

// NewBuffer creates a new buffer.
func NewBuffer() *Buffer {
	return &Buffer{}
}

// Command interface defines the contract for all editor operations.
// It operates on the Buffer (Receiver).
type Command interface {
	Do(b *Buffer)
	Undo(b *Buffer)
}

// Logger defines the interface for logging.
type Logger interface {
	Log(message string)
}
