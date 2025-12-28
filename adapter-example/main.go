package main

import (
	"adapter-example/adapter"
	"adapter-example/domain"
	"fmt"
)

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com domain.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func main() {
	fmt.Println("=== Adapter Pattern ===")

	client := &Client{}
	mac := &adapter.Mac{}

	fmt.Println("\n[Client] Using Mac:")
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsAdapter := &adapter.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	fmt.Println("\n[Client] Using Windows via Adapter:")
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
