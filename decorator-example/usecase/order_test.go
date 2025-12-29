package usecase_test

import (
	"strings"
	"testing"

	"decorator-example/usecase"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

type MockBeverage struct {
	Description string
	Cost        float64
}

func (m *MockBeverage) GetDescription() string {
	return m.Description
}

func (m *MockBeverage) GetCost() float64 {
	return m.Cost
}

func TestOrderService_ProcessOrder(t *testing.T) {
	logger := &MockLogger{}
	service := usecase.NewOrderService(logger)
	beverage := &MockBeverage{Description: "Test Drink", Cost: 1.50}

	service.ProcessOrder(beverage)

	if len(logger.Logs) != 4 {
		t.Errorf("expected 4 log lines, got %d", len(logger.Logs))
	}

	foundDesc := false
	foundCost := false
	for _, log := range logger.Logs {
		if strings.Contains(log, "Test Drink") {
			foundDesc = true
		}
		if strings.Contains(log, "1.50") {
			foundCost = true
		}
	}

	if !foundDesc {
		t.Error("logs missing description")
	}
	if !foundCost {
		t.Error("logs missing cost")
	}
}
