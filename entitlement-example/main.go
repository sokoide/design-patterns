package main

import (
	"entitlement-example/adapter"
	"entitlement-example/infrastructure"
	"entitlement-example/usecase"
	"fmt"
)

func main() {
	// 1. Initialize Infrastructure (Drivers / External Systems)
	adClient := infrastructure.NewPseudoADClient("ldap://dc01.corp.local")
	memoryStore := infrastructure.NewMemoryStore()

	// 2. Initialize Adapters (Component Adapters)
	// These wrap the infrastructure to provide clean internal APIs
	adRepo := adapter.NewADRepo(adClient)
	cacheRepo := adapter.NewCacheRepo(memoryStore)

	// 3. Initialize Gateway (The Main Adapter)
	// This implements the Domain Interface by orchestrating the component adapters
	gateway := adapter.NewInfraGateway(cacheRepo, adRepo)

	// 4. Initialize Usecase
	// Inject the Gateway (which fulfills the domain.EntitlementRepository interface)
	uc := usecase.NewCheckEntitlementUsecase(gateway)

	// 5. Run Scenarios
	fmt.Println("=== Entitlement System Started ===")

	// Scenario 1: Alice in Admin (First time - Cache Miss)
	uc.VerifyUserAccess("admin", "alice")

	// Scenario 2: Alice in Admin (Second time - Cache Hit)
	uc.VerifyUserAccess("admin", "alice")

	// Scenario 3: Bob in Admin (Not a member - Cache Miss)
	uc.VerifyUserAccess("admin", "bob")

	// Scenario 4: Bob in Admin (Second time - Cache Hit for "False")
	// Negative caching is also working
	uc.VerifyUserAccess("admin", "bob")
}
