package main

import (
	"fmt"
	"memento-example/adapter"
)

func main() {
	fmt.Println("=== Memento Pattern ===")

	caretaker := &adapter.Caretaker{}
	editor := &adapter.Editor{}

	editor.Type("This is the first sentence.")
	caretaker.AddMemento(editor.CreateMemento())

	editor.Type("This is the second.")
	caretaker.AddMemento(editor.CreateMemento())

	editor.Type("And this is the third.")
	fmt.Printf("Current Content: %s\n", editor.GetContent())

	editor.Restore(caretaker.GetMemento(1))
	fmt.Printf("Restored to State 2: %s\n", editor.GetContent())

	editor.Restore(caretaker.GetMemento(0))
	fmt.Printf("Restored to State 1: %s\n", editor.GetContent())
}
