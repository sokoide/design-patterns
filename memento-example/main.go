package main

import (
	"fmt"
	"memento-example/adapter"
	"memento-example/usecase"
)

func main() {
	fmt.Println("=== Memento Pattern ===")

	// 1. Setup Dependencies
	editor := &adapter.Editor{}
	logger := adapter.NewConsoleLogger()

	// 2. Setup Caretaker (WriterService)
	service := usecase.NewWriterService(editor, logger)

	// 3. Scenario
	service.Write("This is the first sentence.")
	service.Save()

	service.Write("This is the second.")
	service.Save()

	service.Write("And this is the third.")
	fmt.Printf("Current Content: %s\n", editor.GetContent())

	// Undo to previous saved state
	service.Undo()
	fmt.Printf("Restored to State 2: %s\n", editor.GetContent())

	// Undo again
	service.Undo()
	fmt.Printf("Restored to State 1: %s\n", editor.GetContent())
}
