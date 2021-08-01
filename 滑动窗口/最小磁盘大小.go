package main

import "fmt"

/*
3.某一个大文件被拆成了N个小文件，每个小文件编号从0至N-1，相应大小分别记为S(i)。给定磁盘空间为C，试实现一个函数从N个文件中连续选出
  若干个文件拷贝到磁盘中，使得磁盘剩余空间最小。

函数定义如下：
int MaximumCopy(const std::vector<size_t> &s, size_t C, size_t &start_index, size_t &end_index);
函数返回值为剩余空间，如无解返回-1。

其中start_index, end_index为文件的编号。
如N=5，S = {1, 2, 3, 5, 4}，C = 7
结果为start_index = 0, end_index = 2, return = 1
*/
func main() {
	s := []int{2, 2, 3, 5, 4}
	c := 10 // c := 1
	start_index := -1
	end_index := -1
	r := maximumCopy(s, c, &start_index, &end_index)
	fmt.Println(r, start_index, end_index)
}

func maximumCopy(s []int, c int, start_index, end_index *int) int {
	sl := len(s)
	if sl <= 0 {
		return -1
	}
	max, eIdx, sIdx := 0, -1, -1
	//双指针循环
	for i := 0; i < sl-1; i++ {
		//定义一个参数，接收累加值
		sum := 0
		//移动快指针，使累加值趋近目标值
		for j := i; j < sl-1; j++ {
			//fmt.Printf("i:=%d, j=%d =>>>> \n", i, j)
			sum = sum + s[j]
			if sum <= c {
				//fmt.Printf("i:=%d, j=%d, max=%d，sum=%d, sIdx=%d, eIdx=%d\n", i, j, max, sum, sIdx, eIdx)
				if sum >= max {
					max = sum
					sIdx = i
					eIdx = j
				}
			} else {
				break
			}
		}
	}
	fmt.Printf("max=%d，sIdx=%d, eIdx=%d\n", max, sIdx, eIdx)
	if sIdx == -1 || eIdx == -1 {
		return -1
	}
	*start_index = sIdx
	*end_index = eIdx
	return 1
}
