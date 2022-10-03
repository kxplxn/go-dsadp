package main

import "fmt"

// https://refactoring.guru/design-patterns/decorator

// NOTES:
//
// Although Proxy and Decorator patterns have similar structures, they differ in
// intention; while Proxy's prime purpose is to facilitate ease of use or
// controlled access, a Decorator attaches additional responsibilities. Both
// patterns are built on the composition principle, where one object is supposed
// to delegate some work to another. The difference is that a Proxy usually
// manages the life cycle of its service object on its own, whereas the
// composition of Decorators is always controlled by the client.

// Pizza is the Component interface. It is the common interface for both
// wrappers and wrapped objects.
type Pizza interface {
	getPrice() int
}

// VeggieMania is the Concrete Component. It is an objects being wrapped. It
// defines basic behavior, which can be altered by decorators.
type VeggieMania struct {
}

func (v *VeggieMania) getPrice() int {
	return 15
}

// TomatoTopping is a Concrete Decorator. It defines extra behaviors that can be
// added to components dynamically. Concrete decorators execute their behavior
// either before or after calling the parent method.
type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 7
}

// CheeseTopping is a Concrete Decorator. It defines extra behaviors that can be
// added to components dynamically. Concrete decorators execute their behavior
// either before or after calling the parent method.
type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}

func main() {
	pizza := &VeggieMania{}

	// add cheese topping
	pizzaWithCheese := &CheeseTopping{pizza: pizza}

	// add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf(
		"Price of VeggieMania with tomato and cheese toppings is %d\n",
		pizzaWithCheeseAndTomato.getPrice(),
	)
}
