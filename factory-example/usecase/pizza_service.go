package usecase

import (
	"factory-example/domain"
	"fmt"
)

// PizzaService represents the business logic for accepting pizza orders.
// It depends on a PizzaFactory so that it stays agnostic to the specific pizza type.
type PizzaService struct {
	factory domain.PizzaFactory
}

// NewPizzaService wires the factory into the usecase (dependency injection).
func NewPizzaService(f domain.PizzaFactory) *PizzaService {
	return &PizzaService{
		factory: f,
	}
}

// ServePizza prepares and serves the pizza created by the factory.
func (p *PizzaService) ServePizza() {
	fmt.Println("--- Pizza Order ---")
	pizza := p.factory.CreatePizza()
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Serve()
}
