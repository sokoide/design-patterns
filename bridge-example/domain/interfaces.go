package domain

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(p Printer)
}
