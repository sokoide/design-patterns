package domain

import "fmt"

// PaymentMethod defines the interface that all payment strategies must implement.
// This is the abstraction layer in Clean Architecture.
type PaymentMethod interface {
	Pay(amount float64) error
}

// ShippingMethod defines the interface that shipping strategies must implement.
type ShippingMethod interface {
	Ship(destination string) error
}

// OrderContext holds the data required for a full checkout, including payment and shipping.
type OrderContext struct {
	Amount      float64
	Destination string
}

// Custom Errors
var (
	ErrInvalidAmount      = fmt.Errorf("invalid amount")
	ErrInvalidDestination = fmt.Errorf("invalid destination")
)
