package main

import "fmt"

// Open for extension, closed for modification
// Specification pattern
// You should not keep modifing the interface after implementation because it:
// - is already tested
// - may be clients using it
// - it can end up too large
// If each new feature you need to add a new method/validation to the class
// probably you are breaking the open closed principle

// Example: You need to develop a product selector filter
// the users will be able to choose a product by color or size.

// just some type definition
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

// Wrong example! This way to each new filter type we need to created a new
// method.
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for k, v := range products {
		if v.color == color {
			result = append(result, &products[k])
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for k, v := range products {
		if v.size == size {
			result = append(result, &products[k])
		}
	}

	return result
}

func (f *Filter) FilterByColorSize(products []Product, color Color,
	size Size) []*Product {
	result := make([]*Product, 0)

	for k, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[k])
		}
	}

	return result
}

// Implementing Specification Pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Now each new specification filter is a extension. You can create new
// specifications without change the original interface
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

// composition design pattern to create a composition specification
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Println("Green products: ")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Println(v.name)
	}

	// new implementation
	fmt.Println("Green products with better filter: ")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Println(v.name)
	}

	fmt.Println("Large green products with composition: ")
	largeSpec := SizeSpecification{large}
	andSpec := AndSpecification{greenSpec, largeSpec}
	for _, v := range bf.Filter(products, andSpec) {
		fmt.Println(v.name)
	}
}
