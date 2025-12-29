package domain

import (
	"strings"
	"testing"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(msg string) {
	m.Logs = append(m.Logs, msg)
}

func (m *MockLogger) Contains(substring string) bool {
	for _, l := range m.Logs {
		if strings.Contains(l, substring) {
			return true
		}
	}
	return false
}

func TestCompositePattern(t *testing.T) {
	logger := &MockLogger{}

	// Create files
	file1 := &File{Name: "app.exe", Logger: logger}
	file2 := &File{Name: "config.json", Logger: logger}
	file3 := &File{Name: "notes.txt", Logger: logger}

	// Create directories
	binDir := &Directory{Name: "bin", Logger: logger}
	binDir.Add(file1)

	configDir := &Directory{Name: "config", Logger: logger}
	configDir.Add(file2)

	rootDir := &Directory{Name: "root", Logger: logger}
	rootDir.Add(binDir)
	rootDir.Add(configDir)
	rootDir.Add(file3)

	t.Run("Display recursively", func(t *testing.T) {
		logger.Logs = nil
		rootDir.Display("")

		expectedItems := []string{"root", "bin", "app.exe", "config", "config.json", "notes.txt"}
		for _, item := range expectedItems {
			if !logger.Contains(item) {
				t.Errorf("Expected logs to contain %s, but it didn't", item)
			}
		}
	})

	t.Run("Search recursively", func(t *testing.T) {
		logger.Logs = nil
		rootDir.Search("config")

		if !logger.Contains("Found 'config' in Directory: config") {
			t.Error("Expected to find 'config' in config directory")
		}
		if !logger.Contains("Found 'config' in File: config.json") {
			t.Error("Expected to find 'config' in config.json file")
		}
		if logger.Contains("app.exe") {
			t.Error("Did not expect to find 'config' in app.exe")
		}
	})
}
