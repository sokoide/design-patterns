package main

import (
	"fmt"

	"adapter-example/adapter"
	"adapter-example/usecase"
)

func main() {
	fmt.Println("=== Adapter Pattern ===")

	logger := adapter.NewConsoleLogger()
	client := usecase.NewClient(logger)

	fmt.Println("\n[Client] Using Mac:")
	mac := adapter.NewMac(logger)
	client.InsertLightningConnectorIntoComputer(mac)

	fmt.Println("\n[Client] Using Windows via Adapter:")
	windowsMachine := adapter.NewWindows(logger)
	windowsAdapter := adapter.NewWindowsAdapter(windowsMachine, logger)
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
