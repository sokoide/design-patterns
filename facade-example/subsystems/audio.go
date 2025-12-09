package subsystems

import "fmt"

type AudioSystem struct{}

func NewAudioSystem() *AudioSystem {
	return &AudioSystem{}
}

func (a *AudioSystem) On() {
	fmt.Println("  [Audio] System powered ON.")
}

func (a *AudioSystem) Off() {
	fmt.Println("  [Audio] System powered OFF.")
}

func (a *AudioSystem) SetVolume(vol int) {
	fmt.Printf("  [Audio] Volume set to %d.\n", vol)
}

func (a *AudioSystem) SetSource(source string) {
	fmt.Printf("  [Audio] Source set to '%s'.\n", source)
}

