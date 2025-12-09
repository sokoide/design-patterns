package adapter

import "decorator-example/domain"

// --- Concrete Components (Base Beverages) ---

// Espresso is a concrete component.
type Espresso struct{}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) GetCost() float64 {
	return 1.99
}

// HouseBlend is another concrete component.
type HouseBlend struct{}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{}
}

func (h *HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (h *HouseBlend) GetCost() float64 {
	return 0.89
}

// --- Decorators ---

// Mocha is a Decorator. It wraps a Beverage.
type Mocha struct {
	beverage domain.Beverage
}

func NewMocha(b domain.Beverage) *Mocha {
	return &Mocha{beverage: b}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) GetCost() float64 {
	return m.beverage.GetCost() + 0.20
}

// Whip is another Decorator.
type Whip struct {
	beverage domain.Beverage
}

func NewWhip(b domain.Beverage) *Whip {
	return &Whip{beverage: b}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w *Whip) GetCost() float64 {
	return w.beverage.GetCost() + 0.10
}

// Soy is another Decorator.
type Soy struct {
	beverage domain.Beverage
}

func NewSoy(b domain.Beverage) *Soy {
	return &Soy{beverage: b}
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", Soy"
}

func (s *Soy) GetCost() float64 {
	return s.beverage.GetCost() + 0.15
}
