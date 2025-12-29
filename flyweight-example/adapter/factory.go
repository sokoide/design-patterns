package adapter

import (
	"flyweight-example/domain"
)

// TreeFactory caches flyweight TreeType instances.
type TreeFactory struct {
	treeTypes map[string]*domain.TreeType
}

// NewTreeFactory builds a TreeFactory.
func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*domain.TreeType),
	}
}

// GetTreeType returns a shared TreeType for the provided intrinsic state.
func (f *TreeFactory) GetTreeType(name, color, texture string) *domain.TreeType {
	key := name + "-" + color + "-" + texture
	if t, ok := f.treeTypes[key]; ok {
		return t
	}
	t := domain.NewTreeType(name, color, texture)
	f.treeTypes[key] = t
	return t
}

// TotalTypes returns the number of cached tree types.
func (f *TreeFactory) TotalTypes() int {
	return len(f.treeTypes)
}
