package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var item T
		err := errors.New("no items in queue")
		return item, err
	} else {
		item := q.items[0]
		q.items = q.items[1:]
		return item, nil
	}
}

func main() {
	q := &Queue[int]{}

	for i := 0; i < 3; i++ {
		item := i + 1
		q.Enqueue(item)
		fmt.Printf("enqueued: %d\n", item)
		fmt.Printf("items: %v\n", q.items)
	}

	for i := 0; i < 5; i++ {
		if item, err := q.Dequeue(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("dequeued: %d\n", item)
			fmt.Printf("items: %v\n", q.items)
		}
	}

}
