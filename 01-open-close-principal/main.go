// open close principal
// types are open for extension, closed for modification
// specification pattern
package main

import "fmt"

type Specification interface {
	IsSatisfied(p *Product) bool
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			fmt.Println(v)
			result = append(result, &products[i])
		}
	}
	fmt.Println(result)
	return result
}

type Color int

const (
	red Color = iota + 1
	green
	blue
)

type Size int

const (
	small Size = iota + 1
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

func main() {
	apple := Product{name: "apple", color: green, size: small}
	tree := Product{name: "tree", color: green, size: large}
	house := Product{name: "house", color: blue, size: large}

	product := []Product{apple, tree, house}

	fmt.Printf("Green products:\n")
	greenSpec := ColorSpecification{green}

	bf := BetterFilter{}
	for _, v := range bf.Filter(product, greenSpec) {
		fmt.Printf("- %s is green\n", v.name)
	}
}
