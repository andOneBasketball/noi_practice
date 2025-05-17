package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	i, j := 0, 0
	chCount := make(map[byte]int, len(s))
	for j < len(s) {
		chCount[s[j]]++
		for chCount[s[j]] > 1 && i < j {
			chCount[s[i]]--
			i++
		}
		j++
		if j-i > maxLen {
			maxLen = j - i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
