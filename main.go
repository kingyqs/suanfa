package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

/*
				A
		B        		C
	D		E		F
*/
func main() {
	a := &TreeNode{Val: "A"}
	b := &TreeNode{Val: "B"}
	c := &TreeNode{Val: "C"}
	d := &TreeNode{Val: "D"}
	e := &TreeNode{Val: "E"}
	f := &TreeNode{Val: "F"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	fmt.Println("先序遍历：")
	a.PreOrder()

	fmt.Println("\n中序遍历：")
	a.MidOrder()

	fmt.Println("\n后序遍历：")
	a.PostOrder()

	fmt.Println("\n层序遍历：")
	a.LayOrder()
}

//先序遍历：根->左->右
func (tree *TreeNode) PreOrder() {
	if tree == nil {
		return
	}
	fmt.Print(tree.Val, " ")
	tree.Left.PreOrder()
	tree.Right.PreOrder()
}

//中序遍历: 左->中->右
func (tree *TreeNode) MidOrder() {
	if tree == nil {
		return
	}
	tree.Left.MidOrder()
	fmt.Print(tree.Val, " ")
	tree.Right.MidOrder()
}

//后序遍历: 左->右->中
func (tree *TreeNode) PostOrder() {
	if tree == nil {
		return
	}
	tree.Left.PostOrder()
	tree.Right.PostOrder()
	fmt.Print(tree.Val, " ")
}

//层次遍历：队列来实现
func (tree *TreeNode) LayOrder() {
	if tree == nil {
		return
	}
	queue := new(ArrayQueue)
	queue.Add(tree)
	for queue.len > 0 {
		//出队列
		v, _ := queue.Rem().(*TreeNode)
		fmt.Print(v.Val, " ")
		//左子树有值，入队列
		if v.Left != nil {
			queue.Add(v.Left)
		}

		//右子树有值，入队列
		if v.Right != nil {
			queue.Add(v.Right)
		}
	}
}

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
	queue.len = queue.len + 1
}

func (queue *ArrayQueue) Rem() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.len <= 0 {
		panic("empty")
	}
	out := queue.arr[0]
	queue.arr = queue.arr[1:]
	queue.len = queue.len - 1
	return out
}
