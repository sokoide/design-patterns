package usecase_test

import (
	"fmt"
	"testing"

	"state-example/domain"
	"state-example/usecase"
)

// MockState
type MockState struct {
	NameVal string
}

func (m *MockState) Name() string {
	return m.NameVal
}

func (m *MockState) Handle(action domain.Action) (domain.DoorState, string, error) {
	if action == domain.ActionA {
		return &MockState{NameVal: "NEXT"}, "Transitioning", nil
	}
	return m, "", fmt.Errorf("error")
}

// MockLogger
type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestDoorContext_ExecuteAction(t *testing.T) {
	initialState := &MockState{NameVal: "INITIAL"}
	logger := &MockLogger{}

	// This will fail initially because NewDoorContext doesn't accept logger yet
	door := usecase.NewDoorContext(initialState, logger)

	// Test Initial State
	if door.GetStateName() != "INITIAL" {
		t.Errorf("expected INITIAL, got %s", door.GetStateName())
	}

	// Test Action Success
	door.ExecuteAction(domain.ActionA)
	if door.GetStateName() != "NEXT" {
		t.Errorf("expected NEXT, got %s", door.GetStateName())
	}

	// Check logs (we expect some logs)
	if len(logger.Logs) == 0 {
		t.Error("expected logs, got none")
	}

	// Test Action Error
	door.ExecuteAction(domain.ActionB)
	// State should remain NEXT
	if door.GetStateName() != "NEXT" {
		t.Errorf("expected NEXT after error, got %s", door.GetStateName())
	}
}
