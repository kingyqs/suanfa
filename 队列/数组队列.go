package main

import (
	"fmt"
	"sync"
)

//数组队列
type ArrayQueue struct {
	arr  []interface{}
	len  int
	lock sync.Mutex
}

func (queue *ArrayQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.arr = append(queue.arr, v)
	queue.len = len(queue.arr)
}

func (queue *ArrayQueue) Rem() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.len <= 0 {
		panic("empty")
	}
	out := queue.arr[0]
	queue.arr = queue.arr[1:]
	queue.len = len(queue.arr)
	return out
}

func main() {
	q := new(ArrayQueue)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	fmt.Println(q.arr)
	d := q.Rem()
	fmt.Println(d)
	d = q.Rem()
	fmt.Println(d)
	fmt.Println("arr=", q.arr)
}
