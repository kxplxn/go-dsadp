package main

import "fmt"

type Queue[T any] struct {
	items chan T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items <- item
}

func (q *Queue[T]) Dequeue() T {
	return <-q.items
}

func main() {
	q := &Queue[int]{
		items: make(chan int, 16),
	}

	for i := 0; i < 3; i++ {
		item := i + 1
		q.Enqueue(item)
		fmt.Printf("enqueued: %d\n", item)
	}

	for i := 0; i < 3; i++ {
		item := q.Dequeue()
		fmt.Printf("dequeued: %d\n", item)
	}
}
