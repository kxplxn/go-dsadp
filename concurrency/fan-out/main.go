package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int, name string) {
	for {
		n := rand.Intn(3000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Printf("channel %s <- %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("channel %s -> %d\n", name, n)
	}
}

func fanOut[T any](chA <-chan T, chB, chC chan<- T) {
	for val := range chA {
		if rand.Intn(2) == 1 {
			chB <- val
		} else {
			chC <- val
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanOut(chA, chB, chC)
}
