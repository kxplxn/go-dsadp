package main

// https://www.geeksforgeeks.org/what-is-linked-list/

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *List[T]) First() *Node[T] {
	return l.head
}

func (l *List[T]) Push(value T) {
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
	{
		l := &List[int]{}
		l.Push(1)
		l.Push(2)
		l.Push(3)

		n := l.First()
		for n != nil {
			println(n.value)
			n = n.Next()
		}
	}
	{
		l := &List[string]{}
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
