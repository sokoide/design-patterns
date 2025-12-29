package usecase

import (
	"flyweight-example/adapter"
	"flyweight-example/domain"
)

// Forest manages trees and delegates to the flyweight factory.
type Forest struct {
	trees   []*domain.Tree
	factory *adapter.TreeFactory
	drawer  domain.Drawer
}

// NewForest builds a forest with the provided factory and drawer.
func NewForest(factory *adapter.TreeFactory, drawer domain.Drawer) *Forest {
	return &Forest{
		trees:   make([]*domain.Tree, 0),
		factory: factory,
		drawer:  drawer,
	}
}

// PlantTree creates and stores a tree with shared intrinsic state.
func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	tType := f.factory.GetTreeType(name, color, texture)
	tree := domain.NewTree(x, y, tType)
	f.trees = append(f.trees, tree)
}

// Draw renders all trees in the forest.
func (f *Forest) Draw() {
	for _, t := range f.trees {
		t.Draw(f.drawer)
	}
}

// CountTrees returns the number of planted trees.
func (f *Forest) CountTrees() int {
	return len(f.trees)
}
