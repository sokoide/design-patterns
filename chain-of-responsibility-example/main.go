package main

import (
	"chain-of-responsibility-example/adapter"
	"chain-of-responsibility-example/domain"
	"chain-of-responsibility-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("=== Chain of Responsibility Pattern ===")

	logger := adapter.NewConsoleLogger()

	// 1. Create Handlers (Adapters)
	cashier := adapter.NewCashier(logger)
	doctor := adapter.NewDoctor(logger)
	reception := adapter.NewReception(logger)

	// 2. Set Chain: Reception -> Doctor -> Cashier
	reception.SetNext(doctor)
	doctor.SetNext(cashier)

	// 3. Setup Usecase
	// The usecase knows the "head" of the chain (Reception)
	visitService := usecase.NewPatientVisitService(reception)

	// 4. Execute
	patient := &domain.Patient{Name: "abc"}
	visitService.VisitHospital(patient)
}
