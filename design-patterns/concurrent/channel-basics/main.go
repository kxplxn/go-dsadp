package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://gobyexample.com/channels

func process(ch chan int) {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {
	ch := make(chan int)
	go process(ch)
	fmt.Println("waiting for process")
	n := <-ch
	fmt.Printf("process took %dms\n", n)
}
