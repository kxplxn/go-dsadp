package main

import "fmt"

// ISportsFactory is the Abstract Factory. It declares a set of methods for
// creating each of the Abstract Products.
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("brand unrecognised")
	}
}

// Adidas is a Concrete Factory. Concrete Factories implement the creation
// methods of the abstract factory. Each concrete factory corresponds to a
// specific variant of products and creates only those product variants.
type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "Adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "Adidas",
			size: 14,
		},
	}
}

// Nike is a Concrete Factory. Concrete Factories implement the creation
// methods of the abstract factory. Each concrete factory corresponds to a
// specific variant of products and creates only those product variants.
type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 14,
		},
	}
}

// IShoe is an Abstract Product. It is an interface for a set of distinct but
// related products which make up a product family.
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Shoe provides the common functionality between the IShoe Product Family. This
// is not essential to the Factory Method pattern, but is convenient to have
// when there's an overlap between two types. See how it is used in Concrete
// Products below.
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// AdidasShoe is a Concrete Product. It is a specific implementations of
// the IShoe Abstract Product. Each Abstract Product (IShoe/IShirt) must be
// implemented in all given variants (Adidas/Nike).
type AdidasShoe struct {
	Shoe
}

// NikeShoe is a Concrete Product. It is a specific implementations of
// the IShoe Abstract Product. Each Abstract Product (IShoe/IShirt) must be
// implemented in all given variants (Adidas/Nike).
type NikeShoe struct {
	Shoe
}

// IShirt is an Abstract Product. It is an interface for a set of distinct but
// related products which make up a product family.
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Shirt provides the common functionality between the IShirt Product Family.
// This is not essential to the Factory Method pattern, but is convenient to
// have when there's an overlap between two types. See how it is used in
// Concrete Products below.
type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

// AdidasShirt is a Concrete Product. It is a specific implementations of
// the IShirt Abstract Product. Each Abstract Product (IShoe/IShirt) must be
// implemented in all given variants (Adidas/Nike).
type AdidasShirt struct {
	Shirt
}

// NikeShirt is a Concrete Product. It is a specific implementations of
// the IShirt Abstract Product. Each Abstract Product (IShoe/IShirt) must be
// implemented in all given variants (Adidas/Nike).
type NikeShirt struct {
	Shirt
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Printf("Logo: %s\nSize: %d\n", nikeShoe.getLogo(), nikeShoe.getSize())
	fmt.Printf("Logo: %s\nSize: %d\n", nikeShirt.getLogo(), nikeShirt.getSize())

	fmt.Printf("Logo: %s\nSize: %d\n", adidasShoe.getLogo(), adidasShoe.getSize())
	fmt.Printf("Logo: %s\nSize: %d\n", adidasShirt.getLogo(), adidasShirt.getSize())
}
