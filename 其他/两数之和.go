package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 13
	i, i1 := getTwoSum(nums, target)
	fmt.Println(i, i1)
}

func getTwoSum(nums []int, target int) (int, int) {
	l := len(nums)
	m := map[int]int{}
	for i := 0; i < l; i++ {
		k := target - nums[i]
		if v, ok := m[nums[i]]; ok {
			return v, i
		}
		m[k] = i
	}
	return -1, -1
}
