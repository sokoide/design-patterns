package main

import (
	"fmt"
	"prototype-example/adapter"
	"prototype-example/domain"
)

func main() {
	fmt.Println("=== Prototype Pattern ===")

	file1 := &adapter.File{Name: "File1"}
	file2 := &adapter.File{Name: "File2"}
	file3 := &adapter.File{Name: "File3"}

	folder1 := &adapter.Folder{
		Children: []domain.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &adapter.Folder{
		Children: []domain.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2:")
	folder2.Print("  ")

	fmt.Println("\nCloning Folder2...")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for Clone Folder2:")
	cloneFolder.Print("  ")
}
