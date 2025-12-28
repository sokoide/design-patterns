package domain

import "errors"

// Pizza is the Product interface that every concrete pizza must follow.
// The Usecase layer relies on this interface to interact with the pizza without
// knowing its ingredients or preparation steps.
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Serve()
}

// PizzaFactory is the Creator interface within the Factory Method pattern.
// Concrete factories implement this so that the Usecase can request a pizza
// without depending on a specific pizza type.
type PizzaFactory interface {
	CreatePizza() Pizza
}

// ErrUnknownPizzaType is returned by the Simple Factory when an unsupported
// pizza type is requested.
var ErrUnknownPizzaType = errors.New("unknown pizza type")
