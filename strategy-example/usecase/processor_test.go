package usecase_test

import (
	"errors"
	"testing"

	"strategy-example/domain"
	"strategy-example/usecase"
)

// MockPaymentMethod is a mock implementation of domain.PaymentMethod
type MockPaymentMethod struct {
	PayFunc func(amount float64) error
}

func (m *MockPaymentMethod) Pay(amount float64) error {
	if m.PayFunc != nil {
		return m.PayFunc(amount)
	}
	return nil
}

// MockShippingMethod is a mock implementation of domain.ShippingMethod
type MockShippingMethod struct {
	ShipFunc func(destination string) error
}

func (m *MockShippingMethod) Ship(destination string) error {
	if m.ShipFunc != nil {
		return m.ShipFunc(destination)
	}
	return nil
}

// MockLogger is a mock implementation of domain.Logger
type MockLogger struct {
	LogFunc func(message string)
}

func (m *MockLogger) Log(message string) {
	if m.LogFunc != nil {
		m.LogFunc(message)
	}
}

func TestPaymentProcessor_ProcessOrder(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		destination string
		payErr      error
		shipErr     error
		wantErr     error
	}{
		{
			name:        "Success",
			amount:      100.0,
			destination: "Tokyo",
			payErr:      nil,
			shipErr:     nil,
			wantErr:     nil,
		},
		{
			name:        "Invalid Amount",
			amount:      -10.0,
			destination: "Tokyo",
			payErr:      nil,
			shipErr:     nil,
			wantErr:     domain.ErrInvalidAmount,
		},
		{
			name:        "Invalid Destination",
			amount:      100.0,
			destination: "",
			payErr:      nil,
			shipErr:     nil,
			wantErr:     domain.ErrInvalidDestination,
		},
		{
			name:        "Payment Failure",
			amount:      100.0,
			destination: "Tokyo",
			payErr:      errors.New("payment failed"),
			shipErr:     nil,
			wantErr:     errors.New("payment failed"),
		},
		{
			name:        "Shipping Failure",
			amount:      100.0,
			destination: "Tokyo",
			payErr:      nil,
			shipErr:     errors.New("shipping failed"),
			wantErr:     errors.New("shipping failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPay := &MockPaymentMethod{
				PayFunc: func(amount float64) error {
					return tt.payErr
				},
			}
			mockShip := &MockShippingMethod{
				ShipFunc: func(destination string) error {
					return tt.shipErr
				},
			}
			mockLogger := &MockLogger{
				LogFunc: func(message string) {},
			}

			p := usecase.NewPaymentProcessor(mockPay, mockShip, mockLogger)
			ctx := domain.OrderContext{
				Amount:      tt.amount,
				Destination: tt.destination,
			}

			err := p.ProcessOrder(ctx)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.wantErr)
				} else if err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}