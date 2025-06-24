package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	i, sLen := 0, len(s)
	result := make([]string, 0, sLen/k+1)
	for i < sLen {
		if i+k <= sLen {
			result = append(result, s[i:i+k])
		} else {
			str := s[i:sLen]
			for j := sLen - i; j < k; j++ {
				str += string(fill)
			}
			result = append(result, str)
		}
		i += k
	}
	return result
}

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
