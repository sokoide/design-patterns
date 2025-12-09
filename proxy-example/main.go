package main

import (
	"fmt"
	"proxy-example/adapter"
)

func main() {
	fmt.Println("=== Proxy Pattern ===")

	nginx := adapter.NewNginx()
	appUrl := "/app/status"

	for i := 0; i < 4; i++ {
		code, body := nginx.HandleRequest(appUrl, "GET")
		fmt.Printf("Request %d: Code=%d, Body=%s\n", i+1, code, body)
	}
}