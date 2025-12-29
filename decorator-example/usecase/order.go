package usecase

import (
	"decorator-example/domain"
	"fmt"
)

// OrderService acts as a client that works with the component interface.
type OrderService struct {
	logger domain.Logger
}

func NewOrderService(logger domain.Logger) *OrderService {
	return &OrderService{logger: logger}
}

func (o *OrderService) ProcessOrder(b domain.Beverage) {
	o.logger.Log("--- Processing Order ---")
	o.logger.Log(fmt.Sprintf("Item: %s", b.GetDescription()))
	o.logger.Log(fmt.Sprintf("Total Cost: $%.2f", b.GetCost()))
	o.logger.Log("------------------------")
}
