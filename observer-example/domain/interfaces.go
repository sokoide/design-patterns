package domain

import "errors"

// Observer defines the interface that all listeners must implement.
type Observer interface {
	// OnUpdate is called when the Subject changes.
	OnUpdate(event string)
}

// Subject defines the interface for the object being observed.
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

// Logger defines the interface for logging.
type Logger interface {
	Log(message string)
}

// Common errors
var (
	ErrItemNotFound = errors.New("item not found")
)
