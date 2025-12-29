package adapter_test

import (
	"testing"

	"decorator-example/adapter"
	"decorator-example/domain"
)

func TestBeverageCostAndDescription(t *testing.T) {
	// Base: Espresso ($1.99)
	var b domain.Beverage = adapter.NewEspresso()
	if b.GetCost() != 1.99 {
		t.Errorf("Espresso cost mismatch: got %f, want 1.99", b.GetCost())
	}

	// Add Mocha (+$0.20) -> $2.19
	b = adapter.NewMocha(b)
	if b.GetCost() != 2.19 {
		t.Errorf("Espresso+Mocha cost mismatch: got %f, want 2.19", b.GetCost())
	}
	if b.GetDescription() != "Espresso, Mocha" {
		t.Errorf("Espresso+Mocha desc mismatch: got %s", b.GetDescription())
	}

	// Add Whip (+$0.10) -> $2.29
	b = adapter.NewWhip(b)
	// Float comparison might be tricky, use epsilon if needed, but for simple sums it usually works or use formatted string.
	// Let's use a small margin or just check approx.
	if b.GetCost() < 2.289 || b.GetCost() > 2.291 {
		t.Errorf("Espresso+Mocha+Whip cost mismatch: got %f, want 2.29", b.GetCost())
	}
}
