package domain

import "errors"

// Observer defines the interface that all listeners must implement.
// In Clean Architecture, this is an Output Port (or part of it).
type Observer interface {
	// OnUpdate is called when the Subject changes.
	OnUpdate(event string)
}

// Subject defines the interface for the object being observed.
// (Optional in Go, but good for strict pattern definition)
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

// Common errors
var (
	ErrItemNotFound = errors.New("item not found")
)
