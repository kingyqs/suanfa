package main

import "fmt"

/**
https://www.zhihu.com/roundtable/gaoxiaoshualeetcoden
比如有最经典的
	sliding window模式
	Two pointers模式
	快慢指针模式
	合并intervals模式
	cyclic sort模式
	in-place翻转链表模式
	树上的BFS
	树上的DFS
	双Heaps模式
	subsets模式
	二分法变种
	Top K模式
	多路模式（K-ways）
	0/1背包
	拓扑排序。
*/

func main() {
	//str := "hello hello   kk"
	//l := len(str)
	//begin := 0
	//for i := l - 1; i >= 0; i-- {
	//	if str[i] == ' ' {
	//		break
	//	} else {
	//		begin = i
	//	}
	//}
	//ll := l - begin
	//fmt.Println(ll)
	//
	x := 70100
	sum := 26
	num := 0
	fmt.Println(x / sum)
	for x/sum > 26 {
		sum = sum * 26
		num++
		fmt.Println(x / sum)
	}
	fmt.Println(num)
	sum = sum*26 + num
}
