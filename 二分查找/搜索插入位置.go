package main

import "fmt"

/**
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

示例 4:
输入: nums = [1,3,5,6], target = 0
输出: 0

示例 5:
输入: nums = [1], target = 0
输出: 0
*/

func main() {
	nums := []int{-1, 0, 2, 5, 7, 9, 11}
	target := 11
	idx := searchInsert(nums, target)
	fmt.Println(idx)
}

func searchInsert(nums []int, target int) int {
	left := 0
	ret := len(nums)
	right := ret - 1
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			ret = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ret
}
