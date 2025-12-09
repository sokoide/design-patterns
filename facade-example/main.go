package main

import (
	"facade-example/facade"
	"facade-example/subsystems"
	"fmt"
)

func main() {
	fmt.Println("=== Smart Home System (Facade Pattern) ===")

	// 1. Initialize Subsystems
	// The client (main) still needs to know about these *once* to set them up,
	// or they could be initialized inside the Facade constructor if they are hard dependencies.
	light := subsystems.NewLighting()
	audio := subsystems.NewAudioSystem()
	proj := subsystems.NewProjector()
	screen := subsystems.NewScreen()
	coffee := subsystems.NewCoffeeMaker()

	// 2. Create the Facade
	home := facade.NewSmartHomeFacade(light, audio, proj, screen, coffee)

	// 3. Use the simplified interface
	// The client doesn't need to call dim(), screen.down(), etc. manually.

	// Scenario 1: It's Movie Night
	home.StartMovieMode("The Matrix")

	// ... watching movie ...
	fmt.Println("\n...(2 hours later)...")

	home.EndMovieMode()

	// Scenario 2: Next Morning
	home.GoodMorningMode()
}

