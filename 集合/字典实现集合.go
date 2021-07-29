package main

import (
	"fmt"
	"sync"
)

type MySet struct {
	buf  map[int]struct{}
	size int
	lock sync.Mutex
}

func NewSet(cap int) *MySet {
	data := make(map[int]struct{}, cap)
	return &MySet{
		buf: data,
	}
}

func (s *MySet) Add(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.buf[v] = struct{}{}
	s.size = len(s.buf)
}

func main() {
	fmt.Println("main...")
}
