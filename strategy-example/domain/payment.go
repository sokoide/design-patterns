package domain

import "errors"

// PaymentMethod defines the interface that all payment strategies must implement.
// This is the abstraction layer in Clean Architecture.
type PaymentMethod interface {
	Pay(amount float64) error
}

// ShippingMethod defines the interface that shipping strategies must implement.
type ShippingMethod interface {
	Ship(destination string) error
}

// Logger defines the interface for logging.
type Logger interface {
	Log(message string)
}

// OrderContext holds the data required for a full checkout, including payment and shipping.
type OrderContext struct {
	Amount      float64
	Destination string
}

// Custom Errors
var (
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidDestination = errors.New("invalid destination")
)
