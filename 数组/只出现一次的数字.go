package main

import "fmt"

/**
136. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1

示例2:
输入: [4,1,2,1,2]
输出: 4
*/
func main() {
	//a := 10
	//b := 10
	//c := a ^ b
	//fmt.Println(c)

	nums := []int{4, 1, 2, 1, 2}
	d := singleNumber(nums)
	fmt.Println(d)
}

// 主要利用异或 相同数异或为0，数字与0异或为原来数字原理
func singleNumber(nums []int) int {
	num := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		num ^= nums[i]
	}
	return num
}
