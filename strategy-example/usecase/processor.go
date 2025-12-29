package usecase

import (
	"errors"
	"strategy-example/domain"
)

// PaymentProcessor is the context in the Strategy Pattern.
// It coordinates payment and shipping behaviors.
type PaymentProcessor struct {
	payment  domain.PaymentMethod
	shipping domain.ShippingMethod
	logger   domain.Logger
}

// NewPaymentProcessor creates a new processor with an initial payment and shipping strategy.
func NewPaymentProcessor(payment domain.PaymentMethod, shipping domain.ShippingMethod, logger domain.Logger) *PaymentProcessor {
	return &PaymentProcessor{
		payment:  payment,
		shipping: shipping,
		logger:   logger,
	}
}

// SetPaymentStrategy allows changing the payment strategy at runtime.
func (p *PaymentProcessor) SetPaymentStrategy(payment domain.PaymentMethod) {
	p.payment = payment
}

// SetShippingStrategy allows changing the shipping strategy at runtime.
func (p *PaymentProcessor) SetShippingStrategy(shipping domain.ShippingMethod) {
	p.shipping = shipping
}

// ProcessOrder executes the business logic using the injected strategies.
func (p *PaymentProcessor) ProcessOrder(ctx domain.OrderContext) error {
	if ctx.Amount <= 0 {
		return domain.ErrInvalidAmount
	}

	if ctx.Destination == "" {
		return domain.ErrInvalidDestination
	}

	if p.payment == nil || p.shipping == nil {
		return errors.New("payment or shipping strategy is not set")
	}

	p.logger.Log("--- Starting Payment Process ---")
	// The usecase doesn't know *how* the payment is made, only *that* it is made.
	if err := p.payment.Pay(ctx.Amount); err != nil {
		return err
	}
	p.logger.Log("--- Payment Successful ---")

	p.logger.Log("--- Preparing Shipment ---")
	if err := p.shipping.Ship(ctx.Destination); err != nil {
		return err
	}
	p.logger.Log("--- Shipment Scheduled ---")
	return nil
}
