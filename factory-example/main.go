package main

import (
	"factory-example/adapter"
	"factory-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("=== Pizza Shop (Factory Method) ===")

	fmt.Println("\n[Configuration: Plain]")
	plainFactory := adapter.NewPlainPizzaFactory()
	usecase.NewPizzaService(plainFactory).ServePizza()

	fmt.Println("\n[Configuration: Veggie]")
	veggieFactory := adapter.NewVeggiePizzaFactory()
	usecase.NewPizzaService(veggieFactory).ServePizza()

	fmt.Println("\n[Configuration: Japanese]")
	japaneseFactory := adapter.NewJapanesePizzaFactory()
	usecase.NewPizzaService(japaneseFactory).ServePizza()

	fmt.Println("\n[Configuration: Dynamic (Simple Factory)]")
	pizza, err := adapter.SimplePizzaFactory("japanese")
	if err != nil {
		fmt.Println("❌", err)
		return
	}
	pizza.Serve()
}
