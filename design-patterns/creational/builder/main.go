package main

import "fmt"

// https://refactoring.guru/design-patterns/builder

// IHomeBuilder is the Builder interface. It declares product construction steps
// that are common to all types of concrete builders.
type IHomeBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHome() Home
}

func getBuilder(builderType string) IHomeBuilder {
	if builderType == "apartment" {
		return newApartmentBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// Home is a product. It is the resulting object. Products constructed by
// different builders don’t have to belong to the same class hierarchy or
// interface.
type Home struct {
	windowType string
	doorType   string
	floor      int
}

// The Director class defines the order in which to call construction steps, so
// you can create and reuse specific configurations of products.
type Director struct {
	builder IHomeBuilder
}

func newDirector(b IHomeBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b IHomeBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() Home {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHome()
}

// ApartmentBuilder is a Concrete Builder. It provides a specific implementation
// of the construction steps. Concrete builders may technically produce products
// that don’t follow the common interface.
type ApartmentBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newApartmentBuilder() *ApartmentBuilder {
	return &ApartmentBuilder{}
}

func (b *ApartmentBuilder) setWindowType() {
	b.windowType = "PVC Window"
}

func (b *ApartmentBuilder) setDoorType() {
	b.doorType = "Steel Door"
}

func (b *ApartmentBuilder) setNumFloor() {
	b.floor = 2
}

func (b *ApartmentBuilder) getHome() Home {
	return Home{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// IglooBuilder is another Concrete Builder.
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHome() Home {
	return Home{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// The Client must associate one of the builder objects with the director.
// Usually, it’s done just once, via parameters of the director’s constructor.
// Then the director uses that builder object for all further construction.
// However, there’s an alternative approach for when the client passes the
// builder object to the production method of the director. In this case, you
// can use a different builder each time you produce something with the
// director.
func main() {
	normalBuilder := getBuilder("apartment")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
