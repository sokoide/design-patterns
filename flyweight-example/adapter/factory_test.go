package adapter

import (
	"testing"
)

func TestTreeFactory_GetTreeType(t *testing.T) {
	factory := NewTreeFactory()

	// Request same type twice
	t1 := factory.GetTreeType("Oak", "Green", "Rough")
	t2 := factory.GetTreeType("Oak", "Green", "Rough")

	if t1 != t2 {
		t.Error("expected factory to return the same instance for identical parameters")
	}

	// Request different type
	t3 := factory.GetTreeType("Pine", "Green", "Rough")
	if t1 == t3 {
		t.Error("expected factory to return different instances for different parameters")
	}

	if factory.TotalTypes() != 2 {
		t.Errorf("expected 2 total types, got %d", factory.TotalTypes())
	}
}
