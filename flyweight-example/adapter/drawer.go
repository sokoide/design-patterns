package adapter

import "fmt"

// ConsoleDrawer prints drawing output to stdout.
type ConsoleDrawer struct{}

// NewConsoleDrawer builds a ConsoleDrawer.
func NewConsoleDrawer() *ConsoleDrawer {
	return &ConsoleDrawer{}
}

func (d *ConsoleDrawer) Draw(msg string) {
	fmt.Println(msg)
}
