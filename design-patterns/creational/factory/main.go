package main

import "fmt"

// https://refactoring.guru/design-patterns/factory-method

// IGun is the Product, which declares the interface that is common to all
// objects that can be produced by the creator and its subclasses.
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun encapsulates the common functionality between the Concrete Products. This
// is not essential to the Factory Method pattern, but is convenient to do in
// Go if there's an overlap between two types. See how it is used in Concrete
// Products below.
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// ak47 is a Concrete Products. It is a specific implementations of the Product
// interface.
type ak47 struct {
	Gun
}

// newAK47 is a Concrete Creator that is used by the Factory to return a
// specific Concrete Product.
func newAK47() IGun {
	return &ak47{
		Gun: Gun{
			name:  "AK-47",
			power: 4,
		},
	}

}

// musket is a Concrete Products. It is a specific implementations of the Product
// interface.
type musket struct {
	Gun
}

// newAK47 is a Concrete Creator that is used by the Factory to return a
// specific Concrete Product.
func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}

// getGun is the Factory Method that returns new product objects. It’s important
// that the return type of this method matches the product interface.
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

func main() {
	g1, _ := getGun("ak47")
	g2, _ := getGun("musket")

	fmt.Printf("Gun: %s\n", g1.getName())
	fmt.Printf("Power: %d\n", g1.getPower())

	fmt.Printf("Gun: %s\n", g2.getName())
	fmt.Printf("Power: %d\n", g2.getPower())
}
