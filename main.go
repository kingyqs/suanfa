package main

import (
	"fmt"
)

func main() {
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	Qs(list3, 0, len(list3)-1)
	fmt.Println("test...")
}

func Qs(arr []int, begin, end int) {
	if begin < end {
		left := begin
		right := end
		tmp := 0
		for left != right {
			//取第一个元素为基准值
			tmp = arr[left]
			//从右往左扫描，arr[right]<基准值tmp，则交换，否则继续扫描
			for right > left && arr[right] >= tmp {
				right--
			}
			arr[left] = arr[right]

			//从左往右扫描，arr[left]>基准值tmp， 则交换，否则继续扫描
			for right > left && arr[left] <= tmp {
				left++
			}
			arr[right] = arr[left]
		}

		//扫描一轮 left=right，基准值归位
		arr[right] = tmp

		//相对基准值，继续递归左右两边元素
		Qs(arr, begin, right-1)
		Qs(arr, right+1, end)
	}
}
