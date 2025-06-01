package main

import "fmt"

func invertReverse(str string) string {
	s := []rune(str)
	strLen := len(str)
	i, j := 0, strLen-1
	for i < strLen {
		if str[i] == '0' {
			s[j] = '1'
		} else {
			s[j] = '0'
		}
		i++
		j--
	}
	return string(s)
}

func findKthBit(n int, k int) byte {
	dp := make([]string, n)
	dp[0] = "0"
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + "1" + invertReverse(dp[i-1])
	}
	return dp[n-1][k-1]
}

func main() {
	fmt.Println(findKthBit(3, 1))
	fmt.Println(findKthBit(4, 11))
	fmt.Println(findKthBit(1, 1))
	fmt.Println(findKthBit(2, 3))
}
