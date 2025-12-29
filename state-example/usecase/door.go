package usecase

import (
	"fmt"
	"state-example/domain"
)

// DoorContext represents the door system itself.
// It maintains the current state.
type DoorContext struct {
	currentState domain.DoorState
	logger       domain.Logger
}

func NewDoorContext(initialState domain.DoorState, logger domain.Logger) *DoorContext {
	return &DoorContext{
		currentState: initialState,
		logger:       logger,
	}
}

// ExecuteAction accepts an external input (A or B) and delegates logic to the current state.
func (d *DoorContext) ExecuteAction(action domain.Action) {
	initialStateName := d.currentState.Name()
	nextState, msg, err := d.currentState.Handle(action)

	if err != nil {
		d.logger.Log(fmt.Sprintf("[Input %s] (Current: %-20s) -> Error: %v", action, initialStateName, err))
		return
	}

	d.logger.Log(fmt.Sprintf("[Input %s] (Current: %-20s) -> %s -> New State: %s", action, initialStateName, msg, nextState.Name()))
	d.currentState = nextState
}

func (d *DoorContext) GetStateName() string {
	return d.currentState.Name()
}
