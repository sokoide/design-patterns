package adapter

import (
	"flyweight-example/domain"
)

type TreeFactory struct {
	treeTypes map[string]*domain.TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*domain.TreeType),
	}
}

func (f *TreeFactory) GetTreeType(name, color, texture string) *domain.TreeType {
	key := name + "-" + color + "-" + texture
	if t, ok := f.treeTypes[key]; ok {
		return t
	}
	t := &domain.TreeType{Name: name, Color: color, Texture: texture}
	f.treeTypes[key] = t
	return t
}

func (f *TreeFactory) TotalTypes() int {
	return len(f.treeTypes)
}
