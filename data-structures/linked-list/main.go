package main

// https://www.geeksforgeeks.org/what-is-linked-list/

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) First() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Push(value T) {
	node := &Node[T]{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func main() {
	println("integer list")
	{
		l := &LinkedList[int]{}
		l.Push(1)
		l.Push(2)
		l.Push(3)

		n := l.First()
		for n != nil {
			println(n.value)
			n = n.Next()
		}
	}

	println("string list")
	{
		l := &LinkedList[string]{}
		l.Push("one")
		l.Push("two")
		l.Push("three")

		n := l.First()
		for n != nil {
			println(n.value)
			n = n.Next()
		}
	}
}
