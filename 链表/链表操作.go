package main

import "fmt"

type node struct {
	v    int
	next *node
}

func main() {
	head := node{next: nil}
	n1 := node{v: 1, next: nil}
	head.next = &n1

	n2 := node{v: 2, next: nil}
	n1.next = &n2

	n3 := node{v: 3, next: nil}
	n2.next = &n3

	n4 := node{v: 4, next: nil}
	n3.next = &n4

	//printNode(&head)
	//addNode(&head, 5)
	//printNode(&head)
	insertNode(&head, 1, 6)
	printNode(&head)
}

func printNode(n *node) {
	for n != nil {
		fmt.Println("val:", n.v)
		n = n.next
	}
}

func addNode(n *node, v int) {
	for n.next != nil {
		n = n.next
	}
	n.next = &node{v: v}
}

func insertNode(n *node, k, v int) *node {
	p := &node{v: v}
	for n.next != nil {
		if n.v == k {
			tmp := n.next
			n.next = p
			p.next = tmp
			return n
		}
		n = n.next
	}
	n.next = p
	return n
}
