package main

import "fmt"

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)
}

func SelectSort(arr []int) {
	//扫描数组，找到数组最小值放到最左边，然后依次找第二小
	l := len(arr)
	for i := 0; i < l; i++ {
		min := arr[i]
		minIdx := i
		for j := i + 1; j < l; j++ {
			if arr[j] < min {
				min = arr[j]
				minIdx = j
			}
		}

		//扫描完一次找到最小值，然后进行交换
		if i != minIdx {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
