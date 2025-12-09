package subsystems

import "fmt"

type Projector struct{}

func NewProjector() *Projector {
	return &Projector{}
}

func (p *Projector) On() {
	fmt.Println("  [Projector] Powered ON.")
}

func (p *Projector) Off() {
	fmt.Println("  [Projector] Powered OFF.")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("  [Projector] Input set to '%s'.\n", input)
}

func (p *Projector) WideScreenMode() {
	fmt.Println("  [Projector] Aspect ratio set to 16:9.")
}

type Screen struct{}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Down() {
	fmt.Println("  [Screen] Lowering screen...")
}

func (s *Screen) Up() {
	fmt.Println("  [Screen] Raising screen...")
}

