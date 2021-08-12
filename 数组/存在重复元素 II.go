package main

import "fmt"

/**
219. 存在重复元素 II
给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，使得nums [i] = nums [j]，并且 i 和 j的差的 绝对值 至多为 k。
示例1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/
func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	r := containsNearbyDuplicate(nums, k)
	fmt.Println(r)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]struct{}{}
	n := len(nums)
	//fmt.Println(n)
	for i := 0; i < n; i++ {
		//fmt.Println(">>>> i=", i)
		for j := i + 1; j < n; j++ {
			//fmt.Println("i=", i, ", j=", j, "nums[i]=", nums[i], "nums[j]=", nums[j])
			if nums[i] == nums[j] {
				m[j-i] = struct{}{}
				if _, exit := m[k]; exit {
					return true
				}
			}
		}
	}
	return false
}
