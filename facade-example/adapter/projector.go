package adapter

import (
	"facade-example/domain"
	"fmt"
)

type Projector struct {
	logger domain.Logger
}

func NewProjector(logger domain.Logger) *Projector {
	return &Projector{logger: logger}
}

func (p *Projector) On() {
	p.logger.Log("  [Projector] Powered ON.")
}

func (p *Projector) Off() {
	p.logger.Log("  [Projector] Powered OFF.")
}

func (p *Projector) SetInput(input string) {
	p.logger.Log(fmt.Sprintf("  [Projector] Input set to '%s'.", input))
}

func (p *Projector) WideScreenMode() {
	p.logger.Log("  [Projector] Aspect ratio set to 16:9.")
}
