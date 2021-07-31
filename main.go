package main

import "fmt"

type H struct {
	v int
}

func main() {
	a := &H{v: 10}
	fmt.Println(a)
	a.tt()
	fmt.Println(a)
}

func (h *H) tt() {
	newNode := new(H)
	newNode.v = 12

	tmp := h
	fmt.Println(tmp)
	tmp = (*H)(nil)
	tmp = newNode
}
