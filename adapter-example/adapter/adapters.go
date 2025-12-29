package adapter

import (
	"adapter-example/domain"
)

// Ensure implementations
var (
	_ domain.Computer = (*Mac)(nil)
	_ domain.Computer = (*WindowsAdapter)(nil)
)

// Mac implements Computer interface directly.
type Mac struct {
	logger domain.Logger
}

func NewMac(logger domain.Logger) *Mac {
	return &Mac{logger: logger}
}

func (m *Mac) InsertIntoLightningPort() {
	m.logger.Log("Lightning connector is plugged into Mac machine.")
}

// Windows represents a legacy or incompatible component.
type Windows struct {
	logger domain.Logger
}

func NewWindows(logger domain.Logger) *Windows {
	return &Windows{logger: logger}
}

func (w *Windows) insertIntoUSBPort() {
	w.logger.Log("USB connector is plugged into Windows machine.")
}

// WindowsAdapter makes Windows look like a Computer (Lightning).
type WindowsAdapter struct {
	windowMachine *Windows
	logger        domain.Logger
}

func NewWindowsAdapter(w *Windows, logger domain.Logger) *WindowsAdapter {
	return &WindowsAdapter{
		windowMachine: w,
		logger:        logger,
	}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	w.logger.Log("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
