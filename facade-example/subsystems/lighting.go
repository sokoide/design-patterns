package subsystems

import "fmt"

type Lighting struct{}

func NewLighting() *Lighting {
	return &Lighting{}
}

func (l *Lighting) On() {
	fmt.Println("  [Lighting] Lights turned ON.")
}

func (l *Lighting) Off() {
	fmt.Println("  [Lighting] Lights turned OFF.")
}

func (l *Lighting) Dim(level int) {
	fmt.Printf("  [Lighting] Lights dimmed to %d%%.\n", level)
}
