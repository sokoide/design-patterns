package adapter

import "fmt"

// ConsoleLogger writes logs to stdout.
type ConsoleLogger struct{}

// NewConsoleLogger creates a ConsoleLogger.
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
