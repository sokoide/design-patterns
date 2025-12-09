package main

import (
	"fmt"
	"strategy-example/adapter"
	"strategy-example/domain"
	"strategy-example/usecase"
)

func main() {
	// 1. Initialize Strategies (Adapters)
	creditCard := adapter.NewCreditCardStrategy("1234567812345678", "123")
	paypal := adapter.NewPayPalStrategy("user@example.com")
	bitcoin := adapter.NewBitcoinStrategy("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
	standardShipping := adapter.NewStandardShippingStrategy("Japan Post", 5)
	expressShipping := adapter.NewExpressShippingStrategy("DHL Express")

	// 2. Initialize Usecase with the default strategy (Credit Card)
	processor := usecase.NewPaymentProcessor(creditCard, standardShipping)

	// 3. Execute Business Logic
	fmt.Println("Scenario 1: Customer pays with Credit Card")
	if err := processor.ProcessOrder(domain.OrderContext{
		Amount:      100.50,
		Destination: "Tokyo",
	}); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("\nScenario 2: Customer switches to PayPal")
	// Strategy Pattern allows switching behavior at runtime
	processor.SetPaymentStrategy(paypal)
	processor.SetShippingStrategy(expressShipping)
	if err := processor.ProcessOrder(domain.OrderContext{
		Amount:      50.00,
		Destination: "Osaka",
	}); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("\nScenario 3: Customer uses Crypto")
	processor.SetPaymentStrategy(bitcoin)
	processor.SetShippingStrategy(standardShipping)
	if err := processor.ProcessOrder(domain.OrderContext{
		Amount:      0.05,
		Destination: "Kyoto",
	}); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
