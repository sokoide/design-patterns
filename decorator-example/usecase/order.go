package usecase

import (
	"decorator-example/domain"
	"fmt"
)

// OrderService acts as a client that works with the component interface.
type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) ProcessOrder(b domain.Beverage) {
	fmt.Println("--- Processing Order ---")
	fmt.Printf("Item: %s\n", b.GetDescription())
	fmt.Printf("Total Cost: $%.2f\n", b.GetCost())
	fmt.Println("------------------------")
}
