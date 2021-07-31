package main

import (
	"fmt"
)

type LinkNode struct {
	v    int
	next *LinkNode
}

func main() {
	l1 := &LinkNode{v: 1}
	l2 := &LinkNode{v: 2}
	l1.next = l2
	l3 := &LinkNode{v: 3}
	l2.next = l3
	l4 := &LinkNode{v: 4}
	l3.next = l4

	h2 := reverse1(l1)
	printLinkList(h2)
	//printLinkList(head)
	//fmt.Println("  ")
	//n := reverse(head.next)
	//printLinkList(n)
}

//1、迭代策略
func reverse1(head *LinkNode) *LinkNode {
	pre := (*LinkNode)(nil)
	cur := head
	for cur != nil {
		/*
			链表：1->2->3->4->nil

			第1轮： 1->nil
			第2轮： 2->1->nil
			第3轮： 3->2->1->nil
			第4轮： 4->3->2->1->nil
		*/

		//临时保存链表
		next := cur.next
		cur.next = pre // cur :v->1, next->nil
		fmt.Println("cur=", cur)
		pre = cur // pre :v->1, next->nil
		fmt.Println("pre=", pre)
		cur = next // cur :v->2, next->&l3
	}
	return pre
}

//2.将链表转成数组，反转后再生成链表
func reverse(h *LinkNode) *LinkNode {
	sL := []int{}
	for h != nil {
		sL = append(sL, h.v)
		h = h.next
	}
	j := len(sL) - 1
	head := new(LinkNode)
	for j >= 0 {
		addNode2(head, sL[j])
		j--
	}
	return head
}
func addNode2(n *LinkNode, v int) {
	for n.next != nil {
		n = n.next
	}
	n.next = &LinkNode{v: v}
}

func printLinkList(l *LinkNode) {
	t := l
	for t != nil {
		fmt.Print(t.v, " ")
		t = t.next
	}
}
