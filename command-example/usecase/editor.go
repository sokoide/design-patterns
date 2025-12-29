package usecase

import (
	"command-example/domain"
)

// Editor acts as the Invoker.
// It holds the Receiver (Buffer) and the History (Stack).
type Editor struct {
	buffer *domain.Buffer
	stack  []domain.Command
	logger domain.Logger
}

func NewEditor(logger domain.Logger) *Editor {
	return &Editor{
		buffer: domain.NewBuffer(),
		logger: logger,
	}
}

// Execute performs a command and pushes it to the stack.
func (e *Editor) Execute(c domain.Command) {
	c.Do(e.buffer)
	e.stack = append(e.stack, c)
}

// Undo pops the last command and reverses it.
func (e *Editor) Undo() {
	if len(e.stack) == 0 {
		e.logger.Log("[WARN] Nothing to undo.")
		return
	}

	// Pop
	lastIndex := len(e.stack) - 1
	cmd := e.stack[lastIndex]
	e.stack = e.stack[:lastIndex]

	// Execute Undo logic
	cmd.Undo(e.buffer)
}

// GetContent returns the current buffer content.
func (e *Editor) GetContent() string {
	return e.buffer.Content
}
