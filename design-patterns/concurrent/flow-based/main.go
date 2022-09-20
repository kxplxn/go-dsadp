package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for {
		n := rand.Intn(3000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Printf("producer -> %d\n", n)
		ch <- n
	}
}

func transformer(inCh <-chan int, outCh chan<- string) string {
	for {
		val := <-inCh
		fmt.Printf("transformer <- %d  ", val)
		out := fmt.Sprintf("number %d", val)
		fmt.Printf("-> %s\n", out)
		outCh <- out
	}
}

func consumer(ch <-chan string) {
	for val := range ch {
		fmt.Printf("consumer <- %s\n", val)
	}
}

func block(n int, pdcr func()) {
	for i := 0; i < n; i++ {
		go pdcr()
	}
}

func main() {
	chA := make(chan int, 10)
	chB := make(chan string, 10)

	block(10, func() { producer(chA) })
	block(10, func() { transformer(chA, chB) })
	block(10, func() { consumer(chB) })

	time.Sleep(time.Duration(10000) * time.Millisecond)
}
