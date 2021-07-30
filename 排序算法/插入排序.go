package main

import "fmt"

func main() {
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list2)
	fmt.Println(list2)
}

func InsertSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		deal := arr[i]
		j := i - 1
		if deal < arr[j] {
			for ; j >= 0 && deal < arr[j]; j-- {
				arr[j+1] = arr[j]
			}
			//结束待排序插入进去
			arr[j+1] = deal
		}
	}
}
