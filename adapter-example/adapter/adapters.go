package adapter

import (
	"adapter-example/domain"
)

// Ensure implementations.
var (
	_ domain.Computer = (*Mac)(nil)
	_ domain.Computer = (*WindowsAdapter)(nil)
)

// Mac implements Computer interface directly.
type Mac struct {
	logger domain.Logger
}

// NewMac builds a Mac with the provided logger.
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

// NewWindows builds a Windows machine with the provided logger.
func NewWindows(logger domain.Logger) *Windows {
	return &Windows{logger: logger}
}

func (w *Windows) insertIntoUSBPort() {
	w.logger.Log("USB connector is plugged into Windows machine.")
}

// WindowsAdapter makes Windows look like a Computer (Lightning).
type WindowsAdapter struct {
	windowsMachine *Windows
	logger         domain.Logger
}

// NewWindowsAdapter wraps a Windows machine with a Lightning adapter.
func NewWindowsAdapter(windows *Windows, logger domain.Logger) *WindowsAdapter {
	return &WindowsAdapter{
		windowsMachine: windows,
		logger:         logger,
	}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	w.logger.Log("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}
