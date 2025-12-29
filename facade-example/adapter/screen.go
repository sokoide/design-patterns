package adapter

import (
	"facade-example/domain"
)

// Screen is a concrete screen subsystem.
type Screen struct {
	logger domain.Logger
}

// NewScreen builds a screen adapter.
func NewScreen(logger domain.Logger) *Screen {
	return &Screen{logger: logger}
}

func (s *Screen) Down() {
	s.logger.Log("  [Screen] Lowering screen...")
}

func (s *Screen) Up() {
	s.logger.Log("  [Screen] Raising screen...")
}
