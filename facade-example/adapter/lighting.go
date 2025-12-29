package adapter

import (
	"facade-example/domain"
	"fmt"
)

type Lighting struct {
	logger domain.Logger
}

func NewLighting(logger domain.Logger) *Lighting {
	return &Lighting{logger: logger}
}

func (l *Lighting) On() {
	l.logger.Log("  [Lighting] Lights turned ON.")
}

func (l *Lighting) Off() {
	l.logger.Log("  [Lighting] Lights turned OFF.")
}

func (l *Lighting) Dim(level int) {
	l.logger.Log(fmt.Sprintf("  [Lighting] Lights dimmed to %d%%.", level))
}
