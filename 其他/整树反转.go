package main

import "fmt"

func main() {
	//xia := -123 / 10
	//ge := -123 % 10
	//fmt.Println(xia, ge)
	a := reverseNum(-100)
	fmt.Println(a)
}

func reverseNum(x int) int {
	ret := 0
	for x != 0 {
		d := x % 10
		x = x / 10
		ret = ret*10 + d
	}
	return ret
}
