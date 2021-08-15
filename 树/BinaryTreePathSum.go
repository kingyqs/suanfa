package main

import "fmt"

type TreeNodeV2 struct {
	V     int
	Left  *TreeNodeV2
	Right *TreeNodeV2
}

/**
    1
   / \
  2   4
 / \
2   3

*/
var paths []string
var target = 5

func main() {
	////求全路径
	//root := &TreeNodeV2{
	//	V: 1,
	//	Left: &TreeNodeV2{
	//		V: 2,
	//		Left: &TreeNodeV2{
	//			V: 2,
	//		},
	//		Right: &TreeNodeV2{
	//			V: 3,
	//		},
	//	},
	//	Right: &TreeNodeV2{
	//		V: 4,
	//	},
	//}
	//dfs(root, "")
	//fmt.Printf("paths:%v", paths)

	//求全路径等于目标值路径
	root := &TreeNodeV2{
		V: 1,
		Left: &TreeNodeV2{
			V: 2,
			Left: &TreeNodeV2{
				V: 2,
			},
			Right: &TreeNodeV2{
				V: 3,
			},
		},
		Right: &TreeNodeV2{
			V: 4,
		},
	}
	dfsSum(root, "", 0)
	fmt.Printf("paths:%v", paths)
}

//求全部路径
func dfs(t *TreeNodeV2, p string) {
	if t == nil {
		return
	}
	if p == "" {
		p = fmt.Sprintf("%d", t.V)
	} else {
		p = fmt.Sprintf("%s->%d", p, t.V)
	}
	if t.Left == nil && t.Right == nil {
		//fmt.Printf("p: %s\n", p)
		paths = append(paths, p)
	}
	dfs(t.Left, p)
	dfs(t.Right, p)
}

//求全部路径和等于目标值target
func dfsSum(t *TreeNodeV2, p string, sum int) {
	if t == nil {
		return
	}
	if p == "" {
		p = fmt.Sprintf("%d", t.V)
	} else {
		p = fmt.Sprintf("%s->%d", p, t.V)
	}
	sum = sum + t.V
	if t.Left == nil && t.Right == nil && sum == target {
		fmt.Printf("p: %s\n", p)
		paths = append(paths, p)
	}
	dfsSum(t.Left, p, sum)
	dfsSum(t.Right, p, sum)
}
