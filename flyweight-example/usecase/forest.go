package usecase

import (
	"flyweight-example/adapter"
	"flyweight-example/domain"
)

type Forest struct {
	trees   []*domain.Tree
	factory *adapter.TreeFactory
	drawer  domain.Drawer
}

func NewForest(factory *adapter.TreeFactory, drawer domain.Drawer) *Forest {
	return &Forest{
		trees:   make([]*domain.Tree, 0),
		factory: factory,
		drawer:  drawer,
	}
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	tType := f.factory.GetTreeType(name, color, texture)
	tree := &domain.Tree{
		X:    x,
		Y:    y,
		Type: tType,
	}
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	for _, t := range f.trees {
		t.Draw(f.drawer)
	}
}

func (f *Forest) CountTrees() int {
	return len(f.trees)
}
