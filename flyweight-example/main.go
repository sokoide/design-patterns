package main

import (
	"flyweight-example/adapter"
	"flyweight-example/domain"
	"fmt"
)

func main() {
	fmt.Println("=== Flyweight Pattern ===")

	factory := adapter.GetFactory()

	// Creating a forest
	rees := []*domain.Tree{
		{X: 1, Y: 1, Type: factory.GetTreeType("Oak", "Green", "Rough")},
		{X: 2, Y: 3, Type: factory.GetTreeType("Oak", "Green", "Rough")}, // Reused
		{X: 5, Y: 1, Type: factory.GetTreeType("Pine", "DarkGreen", "Smooth")},
		{X: 6, Y: 6, Type: factory.GetTreeType("Oak", "Green", "Rough")}, // Reused
	}

	for _, t := range trees {
		t.Draw()
	}

	fmt.Printf("\nTotal Tree Objects: %d\n", len(trees))
	fmt.Printf("Total TreeTypes (Flyweights): %d\n", factory.TotalTypes())
}
