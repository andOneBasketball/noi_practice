package main

import (
	"fmt"
)

func isValid(s string) bool {
	leftS := make([]rune, 0, len(s))
	pos := 0
	for _, ss := range s {
		if len(leftS) == 0 && (ss == ')' || ss == ']' || ss == '}') {
			return false
		}
		if ss == '(' || ss == '[' || ss == '{' {
			leftS = append(leftS, ss)
			pos++
			// 过程中切换了思路，这里忘了 continue
			continue
		}
		if ss == ')' && leftS[pos-1] == '(' {
			pos--
			leftS = leftS[0 : len(leftS)-1]
			continue
		}
		if ss == ']' && leftS[pos-1] == '[' {
			pos--
			leftS = leftS[0 : len(leftS)-1]
			continue
		}
		if ss == '}' && leftS[pos-1] == '{' {
			pos--
			leftS = leftS[0 : len(leftS)-1]
			continue
		}
		return false
	}
	// 忽略了该场景，只要 pos 大于0就需要返回 false
	if pos > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
