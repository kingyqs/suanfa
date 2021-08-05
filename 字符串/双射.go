package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abbac"
	str := "dog cat cat dog"
	b := wordPattern(pattern, str)
	fmt.Println(b)
}

func wordPattern(pattern, str string) bool {
	if len(pattern) != len(str) {
		return false
	}
	w2s := map[byte]string{}
	s2w := map[string]byte{}
	strArr := strings.Split(str, " ")
	for i, v := range strArr {
		ch := pattern[i]
		if s2w[v] > 0 && s2w[v] != ch || w2s[ch] != "" && w2s[ch] != v {
			return false
		}
		s2w[v] = ch
		w2s[ch] = v
	}
	return true
}
