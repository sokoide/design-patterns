package main

import (
	"abstract-factory-example/adapter"
	"abstract-factory-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("=== Abstract Factory Pattern ===")

	fmt.Println("\n[Client] Ordering Modern Furniture:")
	modernFactory := &adapter.ModernFactory{}
	service1 := usecase.NewFurnishingService(modernFactory)
	if err := service1.FurnishRoom(); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("\n[Client] Ordering Victorian Furniture:")
	victorianFactory := &adapter.VictorianFactory{}
	service2 := usecase.NewFurnishingService(victorianFactory)
	if err := service2.FurnishRoom(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
