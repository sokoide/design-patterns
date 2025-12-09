package main

import (
	"chain-of-responsibility-example/adapter"
	"chain-of-responsibility-example/domain"
	"fmt"
)

func main() {
	fmt.Println("=== Chain of Responsibility Pattern ===")

	cashier := &adapter.Cashier{}
	
	doctor := &adapter.Doctor{}
	doctor.SetNext(cashier)
	
	reception := &adapter.Reception{}
	reception.SetNext(doctor)

	patient := &domain.Patient{Name: "abc"}
	
	// Chain: Reception -> Doctor -> Cashier
	reception.Execute(patient)
}
