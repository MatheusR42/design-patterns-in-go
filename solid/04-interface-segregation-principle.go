package main

import "fmt"

// You should avoid create too broad interfaces(put many methods), because not
// it if so, will be hard to generalize this interface to other applications

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {
	fmt.Println("Printing...")
}

func (m *MultiFunctionPrinter) Fax(d Document) {
	fmt.Println("Faxing... :O")
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	fmt.Println("Scaning...")
}

// this printer do not support all the features, but we need to implement all
// the methods to be able to use the same interface. This is bad, will be better
// segregate the original interface
type OldFashionedPrinter struct{}

func (o *OldFashionedPrinter) Print(d Document) {
	fmt.Println("Printing...")
}

func (o *OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

func (o *OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

// Now let use the Interface Segregation Principle :)
// so let break out the interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// only creating those new interfaces we solved our original problem. But now we
// have another benefit: if we have a new type that can do both we can implement
// it without change the original code and it will work on both cases. Example:

type Photocopier struct{}

func (p *Photocopier) Print(d Document) {
	fmt.Println("Printing...")
}

func (p *Photocopier) Scan(d Document) {
	fmt.Println("Scaning...")
}

// we can also create a new interface that require only two types:
type MultiFunctionDevice interface {
	Printer
	Scanner
}

func main() {
}
