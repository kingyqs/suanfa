package main

import (
	"fmt"
)

/**
1.给定一个字符串s，试实现函数void remove(char *s, char x)，将原输入字符串s 中与x相等的字符删除。
*/
func main() {
	s := "abcdef"
	x := "acd"
	removeStr(&s, x)
	fmt.Println(s)
}

func removeStr(ts *string, x string) {
	xl := len(x)
	//放入map里面，用于比对
	xm := make(map[string]struct{}, xl)
	for i := 0; i < xl; i++ {
		k := fmt.Sprintf("%c", x[i])
		xm[k] = struct{}{}
	}

	s := *ts
	sl := len(s)
	newStr := ""
	for j := 0; j < sl; j++ {
		jc := fmt.Sprintf("%c", s[j])
		//如果再map里面，则s与x相等元素，删除
		if _, ok := xm[jc]; !ok {
			newStr = newStr + jc
		}
	}
	*ts = newStr
}
