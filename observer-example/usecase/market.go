package usecase

import (
	"fmt"
	"observer-example/domain"
)

// MarketSystem acts as the Concrete Subject.
// It manages the state (Item Price) and the list of Observers.
type MarketSystem struct {
	observers []domain.Observer
	itemName  string
	price     float64
	logger    domain.Logger
}

// NewMarketSystem builds a MarketSystem with an initial price.
func NewMarketSystem(item string, price float64, logger domain.Logger) *MarketSystem {
	return &MarketSystem{
		itemName: item,
		price:    price,
		logger:   logger,
	}
}

// --- Subject Implementation ---

// Register adds an observer to the list.
func (m *MarketSystem) Register(o domain.Observer) {
	m.observers = append(m.observers, o)
}

// Unregister removes an observer from the list.
func (m *MarketSystem) Unregister(o domain.Observer) {
	filtered := m.observers[:0]
	for _, observer := range m.observers {
		if observer != o {
			filtered = append(filtered, observer)
		}
	}
	m.observers = filtered
}

// NotifyAll sends the event to all registered observers.
func (m *MarketSystem) NotifyAll() {
	msg := fmt.Sprintf("Price of '%s' changed to $%.2f", m.itemName, m.price)
	m.logger.Log(fmt.Sprintf("\n--- 📢 Notifying %d observers ---", len(m.observers)))

	for _, observer := range m.observers {
		observer.OnUpdate(msg)
	}
}

// --- Business Logic ---

// UpdatePrice sets a new price and notifies observers.
func (m *MarketSystem) UpdatePrice(newPrice float64) {
	m.logger.Log(fmt.Sprintf("\n[Market] Updating price from $%.2f to $%.2f", m.price, newPrice))
	m.price = newPrice

	// When state changes, notify observers!
	m.NotifyAll()
}
