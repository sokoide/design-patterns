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
	service1.FurnishRoom()

	fmt.Println("\n[Client] Ordering Victorian Furniture:")
	victorianFactory := &adapter.VictorianFactory{}
	service2 := usecase.NewFurnishingService(victorianFactory)
	service2.FurnishRoom()
}

