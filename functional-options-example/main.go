package main

import (
	"fmt"
	"time"

	"functional-options-example/pkg/server"
)

func main() {
	fmt.Println("=== Server Configuration (Functional Options) ===")

	// Case 1: Default Server
	fmt.Println("\n1. Default Server:")
	srv1 := server.New()
	srv1.Start()

	// Case 2: Custom Port and Timeout
	fmt.Println("\n2. Custom Port & Timeout:")
	srv2 := server.New(
		server.WithPort(9090),
		server.WithTimeout(5*time.Second),
	)
	srv2.Start()

	// Case 3: Production Configuration (TLS, Custom Host, High Concurrency)
	fmt.Println("\n3. Production Configuration:")
	srv3 := server.New(
		server.WithHost("api.example.com"),
		server.WithPort(443),
		server.WithTLS(true),
		server.WithMaxConn(1000),
	)
	srv3.Start()
}
