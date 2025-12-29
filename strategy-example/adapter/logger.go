package adapter

import "fmt"

// ConsoleLogger implements domain.Logger and prints to stdout.
type ConsoleLogger struct{}

// NewConsoleLogger creates a ConsoleLogger.
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
