package main

import (
	"decorator-example/adapter"
	"decorator-example/usecase"
	"fmt"
)

func main() {
	logger := adapter.NewConsoleLogger()
	orderService := usecase.NewOrderService(logger)

	fmt.Println("=== Starbuzz Coffee Ordering System ===")

	// 1. Order an Espresso with no condiments
	var beverage1 = adapter.NewEspresso()
	orderService.ProcessOrder(beverage1)

	// 2. Order a HouseBlend with Double Mocha and Whip
	// Pattern: Wrap the object recursively
	var beverage2 = adapter.NewHouseBlend()             // Base
	beverage2Mocha := adapter.NewMocha(beverage2)       // Wrap with Mocha
	beverage2Mocha2 := adapter.NewMocha(beverage2Mocha) // Wrap with Mocha again
	beverage2Final := adapter.NewWhip(beverage2Mocha2)  // Wrap with Whip

	// Note: In Go, since interfaces are implicit, we pass the final wrapped struct
	// but we need to be careful about types if we want to re-assign to a variable of interface type.
	// Here we passed structs directly which implement the interface.
	orderService.ProcessOrder(beverage2Final)

	// 3. Order Espresso with Soy
	var beverage3 = adapter.NewEspresso()
	beverage3Soy := adapter.NewSoy(beverage3)
	orderService.ProcessOrder(beverage3Soy)
}
