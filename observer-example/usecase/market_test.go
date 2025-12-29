package usecase_test

import (
	"strings"
	"testing"

	"observer-example/usecase"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

type MockObserver struct {
	LastEvent string
	Count     int
}

func (m *MockObserver) OnUpdate(event string) {
	m.LastEvent = event
	m.Count++
}

func TestMarketSystem_UpdatePrice(t *testing.T) {
	logger := &MockLogger{}
	market := usecase.NewMarketSystem("TestItem", 100.0, logger)
	obs1 := &MockObserver{}
	obs2 := &MockObserver{}

	market.Register(obs1)
	market.Register(obs2)

	// Update Price
	market.UpdatePrice(150.0)

	// Check observers notified
	if obs1.Count != 1 || obs2.Count != 1 {
		t.Errorf("observers not notified correctly")
	}
	if !strings.Contains(obs1.LastEvent, "150.00") {
		t.Errorf("observer got wrong price: %s", obs1.LastEvent)
	}

	// Unregister
	market.Unregister(obs1)
	market.UpdatePrice(200.0)

	if obs1.Count != 1 {
		t.Error("obs1 should not receive 2nd update")
	}
	if obs2.Count != 2 {
		t.Error("obs2 should receive 2nd update")
	}
}
