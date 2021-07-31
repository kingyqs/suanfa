package main

import (
	"fmt"
	"sync"
)

//链表数据结构
type LinkQueueNode struct {
	V    interface{}
	Next *LinkQueueNode
}

//定义队列数据结构
type LinkQueue struct {
	Root *LinkQueueNode
	size int
	lock sync.Mutex
}

func (queue *LinkQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.Root == nil {
		queue.Root = new(LinkQueueNode)
		queue.Root.V = v
	} else {
		//生成新节点
		newNode := new(LinkQueueNode)
		newNode.V = v
		//赋值临时头指针，找到链表尾插入数据，如果直接移动原始指针会发生变化
		tmpRoot := queue.Root
		for tmpRoot.Next != nil {
			tmpRoot = tmpRoot.Next
		}
		tmpRoot.Next = newNode
	}

	queue.size = queue.size + 1
}

func (queue *LinkQueue) Rem() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//判断链表长度，处理空链表情况
	if queue.size <= 0 {
		panic("empty")
	}
	topNode := queue.Root
	queue.Root = queue.Root.Next
	queue.size = queue.size - 1
	return topNode.V
}

func main() {
	q := new(LinkQueue)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	fmt.Println("q.size=", q.size)
	q.print()

	//d := q.Rem()
	//fmt.Println("d=", d)
}

func (queue *LinkQueue) print() {
	t := queue.Root
	for t != nil {
		fmt.Print(t.V, " ")
		t = t.Next
	}
}
