package adapter

import "fmt"

type ConsoleDrawer struct{}

func NewConsoleDrawer() *ConsoleDrawer {
	return &ConsoleDrawer{}
}

func (d *ConsoleDrawer) Draw(msg string) {
	fmt.Println(msg)
}
