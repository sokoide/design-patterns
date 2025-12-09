package subsystems

import "fmt"

type CoffeeMaker struct{}

func NewCoffeeMaker() *CoffeeMaker {
	return &CoffeeMaker{}
}

func (c *CoffeeMaker) On() {
	fmt.Println("  [Coffee] Heating up...")
}

func (c *CoffeeMaker) Brew() {
	fmt.Println("  [Coffee] Brewing espresso...")
}

func (c *CoffeeMaker) Off() {
	fmt.Println("  [Coffee] Powered OFF.")
}
