package domain

import "errors"

// Action represents the user input (Button A or Button B).
type Action string

const (
	ActionA Action = "A" // Functions as Unlock or Open
	ActionB Action = "B" // Functions as Close or Lock
)

// DoorState defines the interface that all concrete states must implement.
// It takes an action and returns the next state and potentially an error/message.
type DoorState interface {
	Name() string
	Handle(action Action) (DoorState, string, error)
}

// Common errors or messages
var (
	ErrInvalidAction = errors.New("action not valid for current state")
)

type Logger interface {
	Log(message string)
}
