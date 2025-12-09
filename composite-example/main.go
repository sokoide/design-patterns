package main

import (
	"composite-example/adapter"
	"fmt"
)

func main() {
	fmt.Println("=== Composite Pattern ===")

	dev1 := &adapter.Developer{Name: "John"}
	dev2 := &adapter.Developer{Name: "Jane"}
	ஜdes1 := &adapter.Designer{Name: "Mike"}

	manager1 := &adapter.Manager{Name: "Bob"}
	manager1.Add(dev1)
	manager1.Add(des1)

	generalManager := &adapter.Manager{Name: "Alice"}
	generalManager.Add(dev2)
	generalManager.Add(manager1)

	fmt.Println("\nOrganization Chart:")
	generalManager.ShowDetails("  ")
}
