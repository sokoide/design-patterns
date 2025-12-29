package usecase

import "memento-example/domain"

// WriterService is the caretaker managing mementos.
type WriterService struct {
	editor  domain.Editor
	history []*domain.Memento
	logger  domain.Logger
}

// NewWriterService builds a WriterService.
func NewWriterService(editor domain.Editor, logger domain.Logger) *WriterService {
	return &WriterService{
		editor:  editor,
		history: make([]*domain.Memento, 0),
		logger:  logger,
	}
}

// Write appends text to the editor.
func (s *WriterService) Write(text string) {
	s.editor.Type(text)
	s.logger.Log("Typed: " + text)
}

// Save stores the current editor state.
func (s *WriterService) Save() {
	m := s.editor.CreateMemento()
	s.history = append(s.history, m)
	s.logger.Log("State Saved")
}

// Undo restores the editor to the previous saved state.
func (s *WriterService) Undo() {
	if len(s.history) == 0 {
		s.logger.Log("No history to undo")
		return
	}
	// Get last saved state
	lastIndex := len(s.history) - 1
	m := s.history[lastIndex]

	// Remove from history (pop)
	s.history = s.history[:lastIndex]

	s.editor.Restore(m)
	s.logger.Log("Restored to previous state")
}
