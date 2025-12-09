package adapter

import (
	"factory-example/domain"
	"fmt"
)

type pizzaBase struct {
	name string
}

func newPizzaBase(name string) *pizzaBase {
	return &pizzaBase{name: name}
}

func (p *pizzaBase) Prepare() {
	fmt.Printf("Preparing %s...\n", p.name)
}

func (p *pizzaBase) Bake() {
	fmt.Printf("Baking %s...\n", p.name)
}

func (p *pizzaBase) Cut() {
	fmt.Printf("Cutting %s...\n", p.name)
}

func (p *pizzaBase) Serve() {
	fmt.Printf("Serving %s...\n", p.name)
}

// --- Concrete Products ---

type PlainPizza struct {
	*pizzaBase
}

func NewPlainPizza() *PlainPizza {
	return &PlainPizza{
		pizzaBase: newPizzaBase("Plain Pizza"),
	}
}

func (p *PlainPizza) Prepare() {
	p.pizzaBase.Prepare()
	fmt.Println("  Dough, tomato sauce, and mozzarella are ready.")
}

func (p *PlainPizza) Bake() {
	p.pizzaBase.Bake()
	fmt.Println("  Oven heats at 220°C for a crispy crust.")
}

func (p *PlainPizza) Cut() {
	p.pizzaBase.Cut()
	fmt.Println("  Cutting into six classic slices.")
}

func (p *PlainPizza) Serve() {
	p.pizzaBase.Serve()
	fmt.Println("  Served with a sprinkle of basil.")
}

type VeggiePizza struct {
	*pizzaBase
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		pizzaBase: newPizzaBase("Veggie Pizza"),
	}
}

func (v *VeggiePizza) Prepare() {
	v.pizzaBase.Prepare()
	fmt.Println("  Seasonal vegetables layered with bell peppers and olives.")
}

func (v *VeggiePizza) Bake() {
	v.pizzaBase.Bake()
	fmt.Println("  Baked at 210°C to keep vegetables crisp.")
}

func (v *VeggiePizza) Cut() {
	v.pizzaBase.Cut()
	fmt.Println("  Eight slices to share with the table.")
}

func (v *VeggiePizza) Serve() {
	v.pizzaBase.Serve()
	fmt.Println("  Served with a drizzle of olive oil.")
}

type JapanesePizza struct {
	*pizzaBase
}

func NewJapanesePizza() *JapanesePizza {
	return &JapanesePizza{
		pizzaBase: newPizzaBase("Japanese Pizza"),
	}
}

func (j *JapanesePizza) Prepare() {
	j.pizzaBase.Prepare()
	fmt.Println("  Teriyaki chicken, nori, and pickled ginger arranged.")
}

func (j *JapanesePizza) Bake() {
	j.pizzaBase.Bake()
	fmt.Println("  Baked gently at 200°C to caramelize toppings.")
}

func (j *JapanesePizza) Cut() {
	j.pizzaBase.Cut()
	fmt.Println("  Cut into four square pieces to honor the style.")
}

func (j *JapanesePizza) Serve() {
	j.pizzaBase.Serve()
	fmt.Println("  Served with extra pickled ginger on the side.")
}

// --- Factory Method Implementations ---

type PlainPizzaFactory struct{}

func NewPlainPizzaFactory() *PlainPizzaFactory {
	return &PlainPizzaFactory{}
}

func (f *PlainPizzaFactory) CreatePizza() domain.Pizza {
	return NewPlainPizza()
}

type VeggiePizzaFactory struct{}

func NewVeggiePizzaFactory() *VeggiePizzaFactory {
	return &VeggiePizzaFactory{}
}

func (f *VeggiePizzaFactory) CreatePizza() domain.Pizza {
	return NewVeggiePizza()
}

type JapanesePizzaFactory struct{}

func NewJapanesePizzaFactory() *JapanesePizzaFactory {
	return &JapanesePizzaFactory{}
}

func (f *JapanesePizzaFactory) CreatePizza() domain.Pizza {
	return NewJapanesePizza()
}

// --- Simple Factory Function ---

func SimplePizzaFactory(style string) (domain.Pizza, error) {
	switch style {
	case "plain":
		return NewPlainPizza(), nil
	case "veggie":
		return NewVeggiePizza(), nil
	case "japanese":
		return NewJapanesePizza(), nil
	default:
		return nil, domain.ErrUnknownPizzaType
	}
}
