package domain

import "fmt"

// Drawer defines the interface for drawing objects.
type Drawer interface {
	Draw(msg string)
}

// TreeType is the flyweight. It stores intrinsic state (shared data).
type TreeType struct {
	name    string
	color   string
	texture string
}

// NewTreeType creates a TreeType with intrinsic state.
func NewTreeType(name, color, texture string) *TreeType {
	return &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

// Draw performs the operation using extrinsic state (x, y) passed in.
func (t *TreeType) Draw(drawer Drawer, x, y int) {
	// The domain logic formats the output, but the side effect is delegated.
	msg := fmt.Sprintf("Drawing %s tree (%s, %s) at (%d, %d)", t.name, t.color, t.texture, x, y)
	drawer.Draw(msg)
}

// Tree is the Context. It stores extrinsic state (unique data) and a reference to the Flyweight.
type Tree struct {
	x        int
	y        int
	treeType *TreeType
}

// NewTree creates a tree with extrinsic state and a shared TreeType.
func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

func (t *Tree) Draw(drawer Drawer) {
	t.treeType.Draw(drawer, t.x, t.y)
}
