package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Server Configuration (Functional Options) ===")

	// Case 1: Default Server
	fmt.Println("\n1. Default Server:")
	srv1 := NewServer()
	srv1.Start()

	// Case 2: Custom Port and Timeout
	fmt.Println("\n2. Custom Port & Timeout:")
	srv2 := NewServer(
		WithPort(9090),
		WithTimeout(5*time.Second),
	)
	srv2.Start()

	// Case 3: Production Configuration (TLS, Custom Host, High Concurrency)
	fmt.Println("\n3. Production Configuration:")
	srv3 := NewServer(
		WithHost("api.example.com"),
		WithPort(443),
		WithTLS(true),
		WithMaxConn(1000),
	)
	srv3.Start()
}
