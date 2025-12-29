package main

import (
	"fmt"

	"facade-example/adapter"
	"facade-example/usecase"
)

func main() {
	fmt.Println("=== Smart Home System (Facade Pattern) ===")

	logger := adapter.NewConsoleLogger()

	// 1. Initialize Subsystems (Adapters)
	light := adapter.NewLighting(logger)
	audio := adapter.NewAudioSystem(logger)
	proj := adapter.NewProjector(logger)
	screen := adapter.NewScreen(logger)
	coffee := adapter.NewCoffeeMaker(logger)

	// 2. Create the Facade (Usecase)
	home := usecase.NewSmartHomeFacade(light, audio, proj, screen, coffee, logger)

	// 3. Use the simplified interface
	// Scenario 1: It's Movie Night
	home.StartMovieMode("The Matrix")

	// ... watching movie ...
	logger.Log("\n...(2 hours later)...")

	home.EndMovieMode()

	// Scenario 2: Next Morning
	home.GoodMorningMode()
}
