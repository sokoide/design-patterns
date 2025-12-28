package main

import (
	"bridge-example/adapter"
	"fmt"
)

func main() {
	fmt.Println("=== Bridge Pattern ===")

	epson := &adapter.Epson{}
	hp := &adapter.Hp{}

	mac := &adapter.Mac{}
	mac.SetPrinter(epson)
	fmt.Println("\n[System] Mac + Epson:")
	mac.Print()

	mac.SetPrinter(hp)
	fmt.Println("\n[System] Mac + HP:")
	mac.Print()

	win := &adapter.Windows{}
	win.SetPrinter(epson)
	fmt.Println("\n[System] Windows + Epson:")
	win.Print()
}
