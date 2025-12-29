package main

import (
	"fmt"
	"state-example/adapter"
	"state-example/domain"
	"state-example/usecase"
)

func main() {
	// Setup Dependencies
	logger := adapter.NewConsoleLogger()

	// Start with a Locked Door
	initialState := adapter.NewLockedState()
	door := usecase.NewDoorContext(initialState, logger)

	fmt.Println("=== Door State Machine System Started ===")
	fmt.Printf("Initial State: %s\n\n", door.GetStateName())

	// Scenario Steps
	// 1. Try to Close/Lock (B) while Locked -> Should stay Locked
	door.ExecuteAction(domain.ActionB)

	// 2. Unlock (A) -> Should become ClosedUnlocked
	door.ExecuteAction(domain.ActionA)

	// 3. Open (A) -> Should become Open
	door.ExecuteAction(domain.ActionA)

	// 4. Try to Open (A) again -> Should stay Open
	door.ExecuteAction(domain.ActionA)

	// 5. Close (B) -> Should become ClosedUnlocked
	door.ExecuteAction(domain.ActionB)

	// 6. Lock (B) -> Should become Locked
	door.ExecuteAction(domain.ActionB)
}
