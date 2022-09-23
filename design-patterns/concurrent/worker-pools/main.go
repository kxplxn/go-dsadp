package main

import (
	"fmt"
	"math/rand"
	"time"
)

func echo(chIn <-chan int, chOut chan<- int) {
	for n := range chIn {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		chOut <- n
	}
}

func produce(ch chan<- int) {
	for i := 0; true; i++ {
		fmt.Printf("-> send: %d\n", i)
		ch <- i
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)

	// launch a number of workers that listens to the
	// input channel and writes to the output channel
	// since echo takes time, spinning up multiple
	// instances of it improves performance
	for i := 0; i < 100; i++ {
		go echo(chA, chB)
	}

	// start producing values into the input channel
	go produce(chA)

	for n := range chB {
		fmt.Printf("<- recv: %d\n", n)
	}
}
