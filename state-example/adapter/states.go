package adapter

import (
	"state-example/domain"
)

// Ensure implementation
var (
	_ domain.DoorState = (*LockedState)(nil)
	_ domain.DoorState = (*ClosedUnlockedState)(nil)
	_ domain.DoorState = (*OpenState)(nil)
)

// --- 1. Locked State ---
// LockedState represents a locked door.
type LockedState struct{}

// NewLockedState builds a LockedState.
func NewLockedState() *LockedState {
	return &LockedState{}
}

func (s *LockedState) Name() string {
	return "LOCKED 🔒"
}

func (s *LockedState) Handle(action domain.Action) (domain.DoorState, string, error) {
	switch action {
	case domain.ActionA:
		// Unlock -> ClosedUnlocked
		return NewClosedUnlockedState(), "Unlocking door...", nil
	case domain.ActionB:
		// Already Locked
		return s, "Door is already locked.", nil
	}
	return s, "", domain.ErrInvalidAction
}

// --- 2. Closed (Unlocked) State ---
// ClosedUnlockedState represents a closed but unlocked door.
type ClosedUnlockedState struct{}

// NewClosedUnlockedState builds a ClosedUnlockedState.
func NewClosedUnlockedState() *ClosedUnlockedState {
	return &ClosedUnlockedState{}
}

func (s *ClosedUnlockedState) Name() string {
	return "CLOSED (UNLOCKED) 🚪"
}

func (s *ClosedUnlockedState) Handle(action domain.Action) (domain.DoorState, string, error) {
	switch action {
	case domain.ActionA:
		// Open -> OpenState
		return NewOpenState(), "Opening door...", nil
	case domain.ActionB:
		// Lock -> LockedState
		return NewLockedState(), "Locking door...", nil
	}
	return s, "", domain.ErrInvalidAction
}

// --- 3. Open State ---
// OpenState represents an open door.
type OpenState struct{}

// NewOpenState builds an OpenState.
func NewOpenState() *OpenState {
	return &OpenState{}
}

func (s *OpenState) Name() string {
	return "OPEN 💨"
}

func (s *OpenState) Handle(action domain.Action) (domain.DoorState, string, error) {
	switch action {
	case domain.ActionA:
		// Already Open
		return s, "Door is already open.", nil
	case domain.ActionB:
		// Close -> ClosedUnlockedState
		return NewClosedUnlockedState(), "Closing door...", nil
	}
	return s, "", domain.ErrInvalidAction
}
