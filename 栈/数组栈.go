package main

import (
	"fmt"
	"sync"
)

type ArraySack struct {
	buf  []interface{}
	size int
	lock sync.Mutex
}

func main() {
	s1 := &ArraySack{}
	s1.push(2)
	s1.push(3)
	s1.push(4)
	s1.push("www")
	fmt.Println(s1)
	v := s1.pop()
	fmt.Println(v)
	fmt.Println(s1)
}

//入栈
func (s *ArraySack) push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.buf = append(s.buf, v)
	s.size = s.size + 1
}

//出栈
func (s *ArraySack) pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.size == 0 {
		panic("empty")
	}
	v := s.buf[s.size-1]
	s.buf = s.buf[:s.size-1]
	s.size = s.size - 1
	return v
}
