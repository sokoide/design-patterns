package main

import (
	"fmt"

	"command-example/adapter"
	"command-example/usecase"
)

func main() {
	logger := adapter.NewConsoleLogger()
	// Initialize the Invoker (which holds the Receiver internally)
	editor := usecase.NewEditor(logger)

	fmt.Println("=== Command Pattern Editor Demo ===")
	fmt.Printf("Initial Buffer: \"%s\"\n\n", editor.GetContent())

	// 1. Type "Hello"
	cmd1 := adapter.NewInsertCommand("Hello", logger)
	editor.Execute(cmd1)
	printStatus(editor)

	// 2. Type " World"
	cmd2 := adapter.NewInsertCommand(" World", logger)
	editor.Execute(cmd2)
	printStatus(editor)

	// 3. Type "!!!"
	cmd3 := adapter.NewInsertCommand("!!!", logger)
	editor.Execute(cmd3)
	printStatus(editor)

	// 4. Delete last 3 chars ("!!!")
	fmt.Println("\n--- Oops, too excited. Deleting '!!!' ---")
	cmdDelete := adapter.NewDeleteCommand(3, logger)
	editor.Execute(cmdDelete)
	printStatus(editor)

	// 5. Undo the delete (Bring back "!!!")
	fmt.Println("\n--- Wait, I wanted them back! (Undo) ---")
	editor.Undo()
	printStatus(editor)

	// 6. Undo the "!!!" insert
	fmt.Println("\n--- Undo again (Remove '!!!') ---")
	editor.Undo()
	printStatus(editor)

	// 7. Undo the " World" insert
	fmt.Println("\n--- Undo again (Remove ' World') ---")
	editor.Undo()
	printStatus(editor)
}

func printStatus(e *usecase.Editor) {
	fmt.Printf("Current Buffer: \"%s\"\n", e.GetContent())
}
