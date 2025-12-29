package domain

import "fmt"

// Drawer defines the interface for drawing objects.
type Drawer interface {
	Draw(msg string)
}

// TreeType is the Flyweight. It stores intrinsic state (shared data).
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

// Draw performs the operation using extrinsic state (x, y) passed in.
func (t *TreeType) Draw(drawer Drawer, x, y int) {
	// The domain logic formats the output, but the side effect is delegated.
	msg := fmt.Sprintf("Drawing %s tree (%s, %s) at (%d, %d)", t.Name, t.Color, t.Texture, x, y)
	drawer.Draw(msg)
}

// Tree is the Context. It stores extrinsic state (unique data) and a reference to the Flyweight.
type Tree struct {
	X, Y int
	Type *TreeType
}

func (t *Tree) Draw(drawer Drawer) {
	t.Type.Draw(drawer, t.X, t.Y)
}
