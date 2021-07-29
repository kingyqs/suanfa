package main

import (
	"fmt"
	"sync"
)

type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}
type LinkStack struct {
	Root *LinkNode
	size int
	lock sync.Mutex
}

func main() {
	l := &LinkStack{}
	l.LinkStackPush(1)
	l.LinkStackPush(2)
	l.LinkStackPush(3)
	l.LinkStackPrint()
	fmt.Println("pop....")
	l.LinkStackPop()
	l.LinkStackPrint()
}

func (s *LinkStack) LinkStackPush(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	//创建新节点
	newNode := &LinkNode{
		Val: v,
	}

	if s.Root == nil {
		s.Root = newNode
	} else {
		oldRoot := s.Root
		s.Root = newNode
		newNode.Next = oldRoot
	}
	s.size = s.size + 1
}

func (s *LinkStack) LinkStackPop() interface{} {
	s.lock.Lock()
	s.lock.Unlock()
	if s.size == 0 {
		panic("empty")
	}
	topNode := s.Root
	v := topNode.Val

	s.Root = topNode.Next
	s.size = s.size - 1
	return v
}

func (s *LinkStack) LinkStackPrint() {
	tmp := s.Root
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
