package main

import "fmt"

// https://www.geeksforgeeks.org/doubly-linked-list/

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *DoublyLinkedList[T]) First() *Node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) Last() *Node[T] {
	return l.tail
}

// Find is only implemented as an exercise and not a part of the essential
// functionality of doubly linked lists.
func (l *DoublyLinkedList[T]) Find(value T) *Node[T] {
	n := l.First()
	for n != nil {
		if n.value == value {
			return n
		} else {
			n = n.Next()
		}
	}
	return nil
}

func (l *DoublyLinkedList[T]) Push(value T) {
	node := &Node[T]{value: value}
	if l.head == nil {
		l.head = node
	} else {
		node.prev = l.tail
		l.tail.next = node
	}
	l.tail = node
}

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func main() {
	l := &DoublyLinkedList[int]{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	println("traverse forward")
	{
		n := l.First()
		for n != nil {
			println(n.value)
			n = n.Next()
		}
	}

	println("traverse backward")
	{
		n := l.Last()
		for n != nil {
			println(n.value)
			n = n.Prev()
		}
	}

	println("find")
	{
		v := 2
		fmt.Printf("looking for %d\n", v)
		n := l.Find(v)
		fmt.Printf("found %#v\n", n)
	}
}
