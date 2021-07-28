package main

import "fmt"

/**
189. 旋转数组
右移动 k 个位置，其中 k 是非负数。

进阶：
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

*/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	d := rotate2(nums, k)
	fmt.Println(d)
}

func rotate1(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}
	nums1 := nums[k+1:]
	nums2 := nums[:k+1]
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	return nums1
}

func rotate2(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}
	left, right := 0, len(nums)-k
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if right < len(nums) {
			newNums[i] = nums[right]
			right++
			continue
		}
		if left <= k {
			newNums[i] = nums[left]
			left++
			continue
		}
	}

	return newNums
}
