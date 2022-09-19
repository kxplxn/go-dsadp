package main

import (
	"errors"
	"fmt"
)

// https://www.geeksforgeeks.org/stack-data-structure/

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	lenItems := len(s.items)
	if lenItems == 0 {
		var item T
		err := errors.New("no items in stack")
		return item, err
	} else {
		item := s.items[lenItems-1]
		s.items = s.items[:lenItems-1]
		return item, nil
	}
}

func main() {
	s := &Stack[int]{}

	for i := 0; i < 3; i++ {
		item := i + 1
		s.Push(item)
		fmt.Printf("pushed: %d\n", item)
		fmt.Printf("items: %v\n", s.items)
	}

	for i := 0; i < 5; i++ {
		if item, err := s.Pop(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("popped: %d\n", item)
			fmt.Printf("items: %v\n", s.items)
		}
	}
}
