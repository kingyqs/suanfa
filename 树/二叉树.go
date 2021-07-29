package main

import (
	"fmt"
	"sync"
)

/**
 构建一棵树后，我们希望遍历它，有四种遍历方法：
	1、先序遍历：先访问根节点，再访问左子树，最后访问右子树。
	2、后序遍历：先访问左子树，再访问右子树，最后访问根节点。
	3、中序遍历：先访问左子树，再访问根节点，最后访问右子树。
	4、层次遍历：每一层从左到右访问每一个节点。
*/

// 二叉树
type TreeNode struct {
	Data  string    //节点存放数据
	Left  *TreeNode //左子树
	Right *TreeNode //右子树
}

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	//先打印根节点
	fmt.Print(tree.Data, " ")
	//再打印左子树
	PreOrder(tree.Left)
	//再打印右子树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印右子树
	MidOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	//先打印左子树
	PostOrder(tree.Left)
	// 再打印右子树
	PostOrder(tree.Right)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
}

// 层次遍历
type LinkNode struct {
	Value *TreeNode
	Next  *LinkNode
}

//链表队列，先近先出
type LinkQueue struct {
	root *LinkNode  //链表起点
	size int        //队列的元素数量
	lock sync.Mutex //为了并发安全使用的锁
}

//入队列
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		//否则新元素插入链表末尾
		//新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}

	//队中元素数量+1
	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	//顶部元素要出队
	topNode := queue.root
	v := topNode.Value

	//将顶部元素的后继接链上
	queue.root = topNode.Next

	//队中元素数量-1
	queue.size = queue.size - 1

	return v
}

//队列中元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
}
