package adapter

import (
	"bridge-example/domain"
	"fmt"
)

// Concrete Implementor 1
type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by EPSON Printer.")
}

// Concrete Implementor 2
type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by HP Printer.")
}

// Refined Abstraction 1
type Mac struct {
	printer domain.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p domain.Printer) {
	m.printer = p
}

// Refined Abstraction 2
type Windows struct {
	printer domain.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for Windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p domain.Printer) {
	w.printer = p
}
