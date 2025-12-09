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
}

func NewMarketSystem(item string, price float64) *MarketSystem {
	return &MarketSystem{
		observers: make([]domain.Observer, 0),
		itemName:  item,
		price:     price,
	}
}

// --- Subject Implementation ---

func (m *MarketSystem) Register(o domain.Observer) {
	m.observers = append(m.observers, o)
}

func (m *MarketSystem) Unregister(o domain.Observer) {
	var cleanList []domain.Observer
	for _, observer := range m.observers {
		if observer != o {
			cleanList = append(cleanList, observer)
		}
	}
	m.observers = cleanList
}

// NotifyAll sends the event to all registered observers.
func (m *MarketSystem) NotifyAll() {
	msg := fmt.Sprintf("Price of '%s' changed to $%.2f", m.itemName, m.price)
	fmt.Printf("\n--- 📢 Notifying %d observers ---\n", len(m.observers))
	
	for _, observer := range m.observers {
		observer.OnUpdate(msg)
	}
}

// --- Business Logic ---

func (m *MarketSystem) UpdatePrice(newPrice float64) {
	fmt.Printf("\n[Market] Updating price from $%.2f to $%.2f\n", m.price, newPrice)
	m.price = newPrice
	
	// When state changes, notify observers!
	m.NotifyAll()
}
