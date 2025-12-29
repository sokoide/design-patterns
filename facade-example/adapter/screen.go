package adapter

import (
	"facade-example/domain"
)

type Screen struct {
	logger domain.Logger
}

func NewScreen(logger domain.Logger) *Screen {
	return &Screen{logger: logger}
}

func (s *Screen) Down() {
	s.logger.Log("  [Screen] Lowering screen...")
}

func (s *Screen) Up() {
	s.logger.Log("  [Screen] Raising screen...")
}
