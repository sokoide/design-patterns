package main

import (
	"fmt"
	"singleton-example/adapter"
)

func main() {
	fmt.Println("=== Singleton Pattern ===")

	for i := 0; i < 3; i++ {
		fmt.Printf("Call %d: ", i+1)
		db := adapter.GetDatabaseInstance()
		db.Connect()
	}
}
