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
	CardNumber string
	CVV        string
}

func NewCreditCardStrategy(cardNumber, cvv string) *CreditCardStrategy {
	return &CreditCardStrategy{
		CardNumber: cardNumber,
		CVV:        cvv,
	}
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	if c.CardNumber == "" {
		return errors.New("credit card number is empty")
	}
	last4 := c.CardNumber
	if len(last4) > 4 {
		last4 = last4[len(last4)-4:]
	}
	fmt.Printf("Paying $%.2f using Credit Card (Last 4: %s)\n", amount, last4)
	return nil
}

// PayPalStrategy implements the PaymentMethod interface for PayPal payments.
type PayPalStrategy struct {
	Email string
}

func NewPayPalStrategy(email string) *PayPalStrategy {
	return &PayPalStrategy{
		Email: email,
	}
}

func (p *PayPalStrategy) Pay(amount float64) error {
	if p.Email == "" {
		return errors.New("paypal account email is empty")
	}
	fmt.Printf("Paying $%.2f using PayPal (Account: %s)\n", amount, p.Email)
	return nil
}

// BitcoinStrategy implements the PaymentMethod interface for Bitcoin payments.
type BitcoinStrategy struct {
	WalletAddress string
}

func NewBitcoinStrategy(wallet string) *BitcoinStrategy {
	return &BitcoinStrategy{
		WalletAddress: wallet,
	}
}

func (b *BitcoinStrategy) Pay(amount float64) error {
	if b.WalletAddress == "" {
		return errors.New("bitcoin wallet address is empty")
	}
	fmt.Printf("Paying $%.2f using Bitcoin (Wallet: %s)\n", amount, b.WalletAddress)
	return nil
}

// StandardShippingStrategy implements the ShippingMethod interface for regular deliveries.
type StandardShippingStrategy struct {
	Carrier     string
	TransitDays int
}

func NewStandardShippingStrategy(carrier string, transitDays int) *StandardShippingStrategy {
	return &StandardShippingStrategy{
		Carrier:     carrier,
		TransitDays: transitDays,
	}
}

func (s *StandardShippingStrategy) Ship(destination string) error {
	if s.Carrier == "" {
		return errors.New("shipping carrier is empty")
	}
	if s.TransitDays <= 0 {
		return errors.New("invalid transit days")
	}
	fmt.Printf("Scheduling standard %s shipping to %s (ETA: %d days)\n", s.Carrier, destination, s.TransitDays)
	return nil
}

// ExpressShippingStrategy implements the ShippingMethod interface for express deliveries.
type ExpressShippingStrategy struct {
	Carrier string
}

func NewExpressShippingStrategy(carrier string) *ExpressShippingStrategy {
	return &ExpressShippingStrategy{
		Carrier: carrier,
	}
}

func (e *ExpressShippingStrategy) Ship(destination string) error {
	if e.Carrier == "" {
		return errors.New("shipping carrier is empty")
	}
	fmt.Printf("Scheduling express %s shipping to %s (Next day delivery)\n", e.Carrier, destination)
	return nil
}
