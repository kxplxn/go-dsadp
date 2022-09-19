package main

// https://www.geeksforgeeks.org/doubly-linked-list/

type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *DoublyLinkedList[T]) First() *Node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) Last() *Node[T] {
	return l.tail
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
}
