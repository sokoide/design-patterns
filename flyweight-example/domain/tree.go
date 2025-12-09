package domain

import "fmt"

// Flyweight
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("Drawing %s tree (%s, %s) at (%d, %d)\n", t.Name, t.Color, t.Texture, x, y)
}

// Context
type Tree struct {
	X, Y int
	Type *TreeType
}

func (t *Tree) Draw() {
	t.Type.Draw(t.X, t.Y)
}

