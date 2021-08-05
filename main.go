package main

import "fmt"

func main() {
	//str := "hello hello   kk"
	//l := len(str)
	//begin := 0
	//for i := l - 1; i >= 0; i-- {
	//	if str[i] == ' ' {
	//		break
	//	} else {
	//		begin = i
	//	}
	//}
	//ll := l - begin
	//fmt.Println(ll)
	//
	x := 70100
	sum := 26
	num := 0
	fmt.Println(x / sum)
	for x/sum > 26 {
		sum = sum * 26
		num++
		fmt.Println(x / sum)
	}
	fmt.Println(num)
	sum = sum*26 + num
}
