package main

import "fmt"

// https://refactoring.guru/design-patterns/bridge

// Computer is the Abstraction. It defines high-level control logic. It relies
// on the implementation object to do the actual low-level work.
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Mac is a Refined Abstracion. Refined Abstractions provide variants of control
// logic. Like their parent, they work with different implementations via the
// general implementation interface.
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("print requested for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows is a Refined Abstracion. Refined Abstractions provide variants of
// control logic. Like their parent, they work with different implementations
// via the general implementation interface.
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("print requested for win")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Printer is the Implementation interface. It declares functionality that is
// common for all Concrete Implementations. An Abstraction can only communicate
// with an Implementation object via methods that are declared here.
//
// The Abstraction may list the same methods as the implementation, but usually
// the Abstraction declares some complex behaviors that rely on a wide variety
// of primitive operations declared by the Implementation.
type Printer interface {
	PrintFile()
}

// Epson is a Concrete Implementation. It is where the work is actually done.
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("printing via an EPSON printer")
}

// HP is a Concrete Implementation. It is where the work is actually done.
type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("printing via a HP printer")
}

func main() {
	hp := &HP{}
	epson := &Epson{}

	mac := &Mac{}

	mac.SetPrinter(hp)
	mac.Print()
	fmt.Println()

	mac.SetPrinter(epson)
	mac.Print()
	fmt.Println()

	win := &Windows{}

	win.SetPrinter(hp)
	win.Print()
	fmt.Println()

	win.SetPrinter(epson)
	win.Print()
	fmt.Println()
}
