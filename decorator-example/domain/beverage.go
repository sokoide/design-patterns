package domain

// Beverage is the Component interface.
// It defines the behavior that both concrete components and decorators must implement.
type Beverage interface {
	GetDescription() string
	GetCost() float64
}
