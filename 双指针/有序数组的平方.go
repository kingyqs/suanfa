package main

import "fmt"

/**
977. 有序数组的平方
给你一个按 非递减顺序 排序的整数数组 nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

*/
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	d := sortedSquares(nums)
	fmt.Println(d)
}

func sortedSquares(nums []int) []int {
	left := 0
	l := len(nums)
	right := l - 1
	newArr := make([]int, l)

	for i := l - 1; i >= 0; i-- {
		if lv, rv := nums[left]*nums[left], nums[right]*nums[right]; lv > rv {
			newArr[i] = lv
			left++
		} else {
			newArr[i] = rv
			right--
		}
	}

	return newArr
}
