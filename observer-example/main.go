package main

import (
	"fmt"
	"observer-example/adapter"
	"observer-example/usecase"
)

func main() {
	// 1. Create the Subject (Observable)
	// This represents a Bitcoin market tracker
	market := usecase.NewMarketSystem("Bitcoin", 30000.00)

	// 2. Create Observers (Listeners)
	emailClient := adapter.NewEmailNotifier("investor@example.com")
	slackBot := adapter.NewSlackNotifier("crypto-alerts")
	logger := adapter.NewLogNotifier()

	fmt.Println("=== Observer Pattern Market Demo ===")

	// 3. Register Observers
	fmt.Println("Setting up listeners...")
	market.Register(emailClient)
	market.Register(slackBot)
	market.Register(logger)

	// 4. Trigger Event (Price Change)
	// All 3 observers should react
	market.UpdatePrice(32000.00)

	// 5. Unregister an Observer
	// The email user unsubscribes
	fmt.Println("\nUnsubscribing email user...")
	market.Unregister(emailClient)

	// 6. Trigger Another Event
	// Only Slack and Log should react
	market.UpdatePrice(29000.00)
}
