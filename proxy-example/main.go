package main

import (
	"fmt"
	"proxy-example/adapter"
	"proxy-example/usecase"
)

func main() {
	fmt.Println("=== Proxy Pattern ===")

	logger := adapter.NewConsoleLogger()

	// 1. Create the Real Subject (AppServer)
	appServer := adapter.NewAppServer(logger)

	// 2. Create the Proxy (Nginx) wrapping the AppServer
	nginx := adapter.NewNginx(appServer, logger)

	// 3. Client interacts with Nginx
	client := usecase.NewClient(logger)
	appUrl := "/app/status"

	// Send 4 requests (limit is 2)
	client.SendRequests(nginx, appUrl, 4)
}
