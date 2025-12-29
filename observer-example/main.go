package main

import (
	"fmt"
	"observer-example/adapter"
	"observer-example/usecase"
)

func main() {
	logger := adapter.NewConsoleLogger()

	// 1. Create the Subject (Observable)
	market := usecase.NewMarketSystem("Bitcoin", 30000.00, logger)

	// 2. Create Observers (Listeners)
	emailClient := adapter.NewEmailNotifier("investor@example.com", logger)
	slackBot := adapter.NewSlackNotifier("crypto-alerts", logger)
	logObserver := adapter.NewLogNotifier(logger)

	fmt.Println("=== Observer Pattern Market Demo ===")

	// 3. Register Observers
	fmt.Println("Setting up listeners...")
	market.Register(emailClient)
	market.Register(slackBot)
	market.Register(logObserver)

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
