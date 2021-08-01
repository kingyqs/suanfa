package main

import "fmt"

/**
105. 从前序与中序遍历序列构造二叉树
给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
例子；
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/

type TreeNodeV1 struct {
	v     int
	Left  *TreeNodeV1
	Right *TreeNodeV1
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preOrder, inOrder)
	fmt.Println(tree)
}

func buildTree(preorder []int, inorder []int) *TreeNodeV1 {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNodeV1{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
