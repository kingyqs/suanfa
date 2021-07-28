package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	ll := rotateRight(&l, 6)
	printList(ll)
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

/**
head: l1->l2->l3->l4->nil
item： l1->l2->l3->l4->nil

成环：
item->head
l1->l2->l3->l4->l1
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	//计算链表的长度
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	fmt.Println("n:", n, k%n)

	//按环状链表计算有几个环
	add := n - k%n
	if add == n {
		return head
	}

	//将链表变成环
	iter.Next = head
	//循环环状链表，找出移动的数据
	for add > 0 {
		iter = iter.Next
		add--
	}

	ret := iter.Next
	iter.Next = nil
	return ret
}
