package main

import (
	"fmt"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func main() {
	strs := []string{"flower", "flow", "flight"}
	s := maxPreStr(strs)
	fmt.Println(s)
}

func maxPreStr(strs []string) string {
	strlen := len(strs)
	if strlen <= 0 {
		return ""
	} else if strlen == 1 {
		return strs[0]
	}
	maxStr := ""
	s := strs[0]
	for i, v := range s {
		flag := true
		for j := 1; j < strlen; j++ {
			//超出不是公共的
			if i >= len(strs[j]) {
				flag = false
				break
			}
			if int32(strs[j][i]) != v {
				flag = false
				break
			}
		}
		//相等是公共
		if flag {
			maxStr = maxStr + fmt.Sprintf("%c", v)
		} else {
			break
		}
	}
	return maxStr
}
