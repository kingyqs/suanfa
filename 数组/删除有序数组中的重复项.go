package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n)
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	slow := 1
	newNums := []int{nums[0]}
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			newNums = append(newNums, nums[fast])
			slow++
		}
	}
	fmt.Println(newNums)
	return slow
}
