package main

import (
	"flyweight-example/adapter"
	"flyweight-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("=== Flyweight Pattern ===")

	factory := adapter.NewTreeFactory()
	drawer := adapter.NewConsoleDrawer()
	forest := usecase.NewForest(factory, drawer)

	// Creating a forest
	forest.PlantTree(1, 1, "Oak", "Green", "Rough")
	forest.PlantTree(2, 3, "Oak", "Green", "Rough") // Reused
	forest.PlantTree(5, 1, "Pine", "DarkGreen", "Smooth")
	forest.PlantTree(6, 6, "Oak", "Green", "Rough") // Reused

	forest.Draw()

	fmt.Printf("\nTotal Tree Objects: %d\n", forest.CountTrees())
	fmt.Printf("Total TreeTypes (Flyweights): %d\n", factory.TotalTypes())
}
