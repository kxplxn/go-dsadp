package main

import "fmt"

// https://www.geeksforgeeks.org/binary-tree-data-structure/

type Tree struct {
	node *Node
}

func (t *Tree) Insert(value int) {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.Insert(value)
	}
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.Insert(value)
		}
	}
}

func (n *Node) Exists(value int) bool {
	if n == nil {
		return false
	} else if value == n.value {
		return true
	} else if value < n.value {
		return n.left.Exists(value)
	} else {
		return n.right.Exists(value)
	}
}

func main() {
	t := &Tree{}

	for _, value := range []int{10, 8, 9, 12, 11, 13} {
		t.Insert(value)
	}
	fmt.Println("all values:")
	printNode(t.node)

	fmt.Println("exists:")
	for _, lookupValue := range []int{10, 17, 93, 11, 7, 8} {
		fmt.Printf("%d %v\n", lookupValue, t.node.Exists(lookupValue))
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)

	printNode(n.left)
	printNode(n.right)
}
