package usecase_test

import (
	"strings"
	"testing"

	"facade-example/usecase"
)

// MockLogger
type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

// Mocks for subsystems
type MockLighting struct{}

func (m *MockLighting) On()           {}
func (m *MockLighting) Off()          {}
func (m *MockLighting) Dim(level int) {}

type MockAudio struct{}

func (m *MockAudio) On()                     {}
func (m *MockAudio) Off()                    {}
func (m *MockAudio) SetVolume(vol int)       {}
func (m *MockAudio) SetSource(source string) {}

type MockProjector struct{}

func (m *MockProjector) On()                   {}
func (m *MockProjector) Off()                  {}
func (m *MockProjector) SetInput(input string) {}
func (m *MockProjector) WideScreenMode()       {}

type MockScreen struct{}

func (m *MockScreen) Up()   {}
func (m *MockScreen) Down() {}

type MockCoffee struct{}

func (m *MockCoffee) On()   {}
func (m *MockCoffee) Off()  {}
func (m *MockCoffee) Brew() {}

func TestSmartHomeFacade_StartMovieMode(t *testing.T) {
	logger := &MockLogger{}
	facade := usecase.NewSmartHomeFacade(
		&MockLighting{},
		&MockAudio{},
		&MockProjector{},
		&MockScreen{},
		&MockCoffee{},
		logger,
	)

	facade.StartMovieMode("Test Movie")

	found := false
	for _, log := range logger.Logs {
		if strings.Contains(log, "Test Movie") {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected movie name in logs")
	}
}
