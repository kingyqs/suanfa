package main

import "fmt"

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)
	fmt.Println(list)
}

//冒泡排序
func BubbleSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		//这里已经把最大交换到最后一位，实际循环交换的次数是迭减
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
