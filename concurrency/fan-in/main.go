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

func fanIn[T any](chA, chB <-chan T, chC chan<- T) {
	for {
		select {
		case val := <-chA:
			chC <- val
		case val := <-chB:
			chC <- val
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC, "C")

	fanIn(chA, chB, chC)
}
