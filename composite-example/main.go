package main

import (
	"fmt"
	"github.com/sokoide/design-patterns/composite-example/domain"
)

// ConsoleLogger implements domain.Logger interface
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func main() {
	logger := &ConsoleLogger{}

	fmt.Println("=== Composite Pattern: File System Demo ===")
	fmt.Println()

	// 1. Setup structure
	// root/
	//   bin/
	//     app.exe
	//     lib.dll
	//   config/
	//     settings.json
	//   readme.txt

	root := &domain.Directory{Name: "root", Logger: logger}
	bin := &domain.Directory{Name: "bin", Logger: logger}
	config := &domain.Directory{Name: "config", Logger: logger}

	file1 := &domain.File{Name: "app.exe", Logger: logger}
	file2 := &domain.File{Name: "lib.dll", Logger: logger}
	file3 := &domain.File{Name: "settings.json", Logger: logger}
	file4 := &domain.File{Name: "readme.txt", Logger: logger}

	bin.Add(file1)
	bin.Add(file2)
	config.Add(file3)

	root.Add(bin)
	root.Add(config)
	root.Add(file4)

	// 2. Display entire structure
	fmt.Println("--- Full Directory Structure ---")
	root.Display("")
	fmt.Println()

	// 3. Search for 'app'
	fmt.Println("--- Search results for 'app' ---")
	root.Search("app")
	fmt.Println()

	// 4. Search for 'json'
	fmt.Println("--- Search results for 'json' ---")
	root.Search("json")
	fmt.Println()

	// 5. Demonstrate transparency: search specifically on 'bin' directory
	fmt.Println("--- Search results for 'lib' within 'bin' directory ---")
	bin.Search("lib")
}
