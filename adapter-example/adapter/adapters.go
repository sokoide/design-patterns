package adapter

import (
	"adapter-example/domain"
	"fmt"
)

// Client: Mac
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into Mac machine.")
}

// Service: Windows (Unknown to Client)
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into Windows machine.")
}

// Adapter
type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.insertIntoUSBPort()
}
