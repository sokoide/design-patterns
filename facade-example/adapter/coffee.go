package adapter

import (
	"facade-example/domain"
)

type CoffeeMaker struct {
	logger domain.Logger
}

func NewCoffeeMaker(logger domain.Logger) *CoffeeMaker {
	return &CoffeeMaker{logger: logger}
}

func (c *CoffeeMaker) On() {
	c.logger.Log("  [Coffee] Heating up...")
}

func (c *CoffeeMaker) Brew() {
	c.logger.Log("  [Coffee] Brewing espresso...")
}

func (c *CoffeeMaker) Off() {
	c.logger.Log("  [Coffee] Powered OFF.")
}
