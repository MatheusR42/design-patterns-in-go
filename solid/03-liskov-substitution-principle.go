package main

import "fmt"

// Liskov Substitution Principle says that we should be able to substitute the
// parent class with a child class without change the program. Example:
// we can have a parent class that has general payments methods (e.g: getValue,
// getStatus, getPayerId). The child classes may be more especialized to handle
// credit or debit cards payments. But those methods should respect the original
// ideal of parent class.
// Liskov Substitution Principle does not apply directly in Go because
// because go do not have inheritance. But there is a situation where we can
// think about it in go. Let see a example with geometric objects

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// this is the original implementation
type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Squared is breaking the Liskov Substitution Principle, because we changed the
// implementation of SetWidth and SetHeight that was created previously.
type Square struct {
	width, height int
}

func (s *Square) GetWidth() int {
	return s.width
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) GetHeight() int {
	return s.height
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("Expected an area of %d, and got %d\n", expectedArea, actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	// square returns an unexpected result compared with the original
	// implementation. It is breaking the Liskov Substitution Principle
	s := &Square{2, 2}
	UseIt(s)
}
