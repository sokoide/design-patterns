package usecase

import (
	"fmt"
	"state-example/domain"
)

// DoorContext represents the door system itself.
// It maintains the current state.
type DoorContext struct {
	currentState domain.DoorState
}

func NewDoorContext(initialState domain.DoorState) *DoorContext {
	return &DoorContext{
		currentState: initialState,
	}
}

// ExecuteAction accepts an external input (A or B) and delegates logic to the current state.
func (d *DoorContext) ExecuteAction(action domain.Action) {
	fmt.Printf("[Input %s] (Current: %-20s) -> ", action, d.currentState.Name())

	nextState, msg, err := d.currentState.Handle(action)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s -> New State: %s\n", msg, nextState.Name())
	d.currentState = nextState
}

func (d *DoorContext) GetStateName() string {
	return d.currentState.Name()
}
