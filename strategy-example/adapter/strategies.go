package adapter

import (
	"errors"
	"fmt"
	"strategy-example/domain"
)

// Ensure that the strategies implement the relevant interfaces.
var (
	_ domain.PaymentMethod  = (*CreditCardStrategy)(nil)
	_ domain.PaymentMethod  = (*PayPalStrategy)(nil)
	_ domain.PaymentMethod  = (*BitcoinStrategy)(nil)
	_ domain.ShippingMethod = (*StandardShippingStrategy)(nil)
	_ domain.ShippingMethod = (*ExpressShippingStrategy)(nil)
)

// CreditCardStrategy implements the PaymentMethod interface for Credit Card payments.
type CreditCardStrategy struct {
	cardNumber string
	cvv        string
}

// NewCreditCardStrategy builds a CreditCardStrategy.
func NewCreditCardStrategy(cardNumber, cvv string) *CreditCardStrategy {
	return &CreditCardStrategy{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	if c.cardNumber == "" {
		return errors.New("credit card number is empty")
	}
	last4 := c.cardNumber
	if len(last4) > 4 {
		last4 = last4[len(last4)-4:]
	}
	fmt.Printf("Paying $%.2f using Credit Card (Last 4: %s)\n", amount, last4)
	return nil
}

// PayPalStrategy implements the PaymentMethod interface for PayPal payments.
type PayPalStrategy struct {
	email string
}

// NewPayPalStrategy builds a PayPalStrategy.
func NewPayPalStrategy(email string) *PayPalStrategy {
	return &PayPalStrategy{
		email: email,
	}
}

func (p *PayPalStrategy) Pay(amount float64) error {
	if p.email == "" {
		return errors.New("paypal account email is empty")
	}
	fmt.Printf("Paying $%.2f using PayPal (Account: %s)\n", amount, p.email)
	return nil
}

// BitcoinStrategy implements the PaymentMethod interface for Bitcoin payments.
type BitcoinStrategy struct {
	walletAddress string
}

// NewBitcoinStrategy builds a BitcoinStrategy.
func NewBitcoinStrategy(wallet string) *BitcoinStrategy {
	return &BitcoinStrategy{
		walletAddress: wallet,
	}
}

func (b *BitcoinStrategy) Pay(amount float64) error {
	if b.walletAddress == "" {
		return errors.New("bitcoin wallet address is empty")
	}
	fmt.Printf("Paying $%.2f using Bitcoin (Wallet: %s)\n", amount, b.walletAddress)
	return nil
}

// StandardShippingStrategy implements the ShippingMethod interface for regular deliveries.
type StandardShippingStrategy struct {
	carrier     string
	transitDays int
}

// NewStandardShippingStrategy builds a StandardShippingStrategy.
func NewStandardShippingStrategy(carrier string, transitDays int) *StandardShippingStrategy {
	return &StandardShippingStrategy{
		carrier:     carrier,
		transitDays: transitDays,
	}
}

func (s *StandardShippingStrategy) Ship(destination string) error {
	if s.carrier == "" {
		return errors.New("shipping carrier is empty")
	}
	if s.transitDays <= 0 {
		return errors.New("invalid transit days")
	}
	fmt.Printf("Scheduling standard %s shipping to %s (ETA: %d days)\n", s.carrier, destination, s.transitDays)
	return nil
}

// ExpressShippingStrategy implements the ShippingMethod interface for express deliveries.
type ExpressShippingStrategy struct {
	carrier string
}

// NewExpressShippingStrategy builds an ExpressShippingStrategy.
func NewExpressShippingStrategy(carrier string) *ExpressShippingStrategy {
	return &ExpressShippingStrategy{
		carrier: carrier,
	}
}

func (e *ExpressShippingStrategy) Ship(destination string) error {
	if e.carrier == "" {
		return errors.New("shipping carrier is empty")
	}
	fmt.Printf("Scheduling express %s shipping to %s (Next day delivery)\n", e.carrier, destination)
	return nil
}
