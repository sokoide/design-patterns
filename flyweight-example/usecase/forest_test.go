package usecase

import (
	"testing"

	"flyweight-example/adapter"
)

type MockDrawer struct {
	DrawCount int
}

func (m *MockDrawer) Draw(msg string) {
	m.DrawCount++
}

func TestForest(t *testing.T) {
	factory := adapter.NewTreeFactory()
	drawer := &MockDrawer{}
	forest := NewForest(factory, drawer)

	forest.PlantTree(1, 1, "Oak", "Green", "Rough")
	forest.PlantTree(2, 2, "Oak", "Green", "Rough")
	forest.PlantTree(3, 3, "Pine", "Green", "Smooth")

	if forest.CountTrees() != 3 {
		t.Errorf("expected 3 trees, got %d", forest.CountTrees())
	}

	forest.Draw()

	if drawer.DrawCount != 3 {
		t.Errorf("expected 3 draw calls, got %d", drawer.DrawCount)
	}
}
