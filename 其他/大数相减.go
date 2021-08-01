package main

import "fmt"

/**
2.考虑有2个存储在string中的十进制正整数（远大于int64）A和B，
现请你实现一个函数std::string minus_hp(const std::string &A, const std::string &B)，计算A-B的结果，
如：
  11111111111111111111111
- 11111111111111111111110
                      = 1
*/

func main() {
	A := "11244"
	B := "1335"
	fu := ""
	la := len(A)
	lb := len(B)
	change := false
	if lb > la {
		change = true
	}
	if la == lb {
		for i := 0; i < la; i++ {
			if B[i] > A[i] {
				change = true
				break
			}
		}
	}
	if change {
		fu = "-"
		A, B = B, A
	}

	r := fu + minusHp(A, B)
	fmt.Println("r=", r)
}

// A >= B
func minusHp(A, B string) string {
	aIdx := len(A) - 1
	bIdx := len(B) - 1
	//计算结果保存值
	ret := ""
	//借位保存值
	jw := 0

	for aIdx >= 0 && bIdx >= 0 {
		//减少掉借位的值
		aTmp := A[aIdx] - uint8(jw)
		bTmp := B[bIdx]
		//判断是否需要向高位借位，如果减值为正值
		if aTmp >= '0' && aTmp >= bTmp {
			//重置借位
			jw = 0
			r := aTmp - bTmp
			ret = fmt.Sprintf("%v", r) + ret

		} else {
			jw = 1
			//如果被减数的值小，则向前位借1，转则位10
			if aTmp >= '0' {
				r := (aTmp + 10) - bTmp
				ret = fmt.Sprintf("%v", r) + ret
			} else {
				aTmp = '0' + 9
				r := aTmp - bTmp
				ret = fmt.Sprintf("%v", r) + ret
			}
		}
		aIdx--
		bIdx--
	}
	for aIdx >= 0 {
		aTmp := A[aIdx] - uint8(jw)
		if aTmp > '0' {
			ret = fmt.Sprintf("%v", A[aIdx]-'0') + ret
		}
		aIdx--
	}
	return ret
}
