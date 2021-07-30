package main

import "fmt"

func main() {
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//QuickSort(list3, 0, len(list3)-1)
	QuickSort2(list3, 0, len(list3)-1)
	fmt.Println(list3)
}

//时间复杂度 O(nlogn)

//递归形式1 参考：https://www.cnblogs.com/MOBIN/p/4681369.html
func QuickSort2(array []int, begin, end int) {
	left := begin
	right := end
	tmp := 0
	if left < right { //待排序的只是有两种情况
		//1.选数组第一个元素为基准值
		tmp = array[left]

		//2.从左右两边交替扫描，直到left=right
		for left != right {
			//3.从右往左扫描，找到第一比基准元素小的元素，并将找到这种元素arr[right]后与arr[left]交换；如果扫描到的值比基准值大，则继续向左扫描
			for right > left && array[right] >= tmp {
				right--
			}
			array[left] = array[right]

			//4.从左往左扫描，找到第一比基准元素大的元素，并将找到这种元素arr[right]后与arr[left]交换；如果扫描到的值比基准值小，则继续向右扫描
			for right > left && array[left] <= tmp {
				left++
			}
			array[right] = array[left]
		}

		//跳出循环后，left == right
		//5.基准值归位
		array[right] = tmp
		//6.对左右两部元素进行递归处理
		QuickSort2(array, begin, left-1) //对基准元素左边的元素进行递归排序
		QuickSort2(array, right+1, end)  //对基准元素右边的元素进行递归排序
	}
}

//递归形式2
func QuickSort(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort(array, loc+1, end)
	}
}
func partition(array []int, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	//没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] //交换
			j--
		} else {
			i++
		}
	}

	/**
	 *	跳出while循环后，i = j。
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                       -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}
	array[begin], array[i] = array[i], array[begin]

	return i
}

//非递归，自定义一个栈，不会存在递归程序栈溢出问题
