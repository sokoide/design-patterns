package adapter

import (
	"facade-example/domain"
	"fmt"
)

type AudioSystem struct {
	logger domain.Logger
}

func NewAudioSystem(logger domain.Logger) *AudioSystem {
	return &AudioSystem{logger: logger}
}

func (a *AudioSystem) On() {
	a.logger.Log("  [Audio] System powered ON.")
}

func (a *AudioSystem) Off() {
	a.logger.Log("  [Audio] System powered OFF.")
}

func (a *AudioSystem) SetVolume(vol int) {
	a.logger.Log(fmt.Sprintf("  [Audio] Volume set to %d.", vol))
}

func (a *AudioSystem) SetSource(source string) {
	a.logger.Log(fmt.Sprintf("  [Audio] Source set to '%s'.", source))
}
