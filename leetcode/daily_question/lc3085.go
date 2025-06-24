package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeletions(word string, k int) int {
	wordLen := len(word)
	m := make(map[byte]int, wordLen)
	for i := 0; i < wordLen; i++ {
		m[word[i]]++
	}
	//fmt.Println(m)
	count := make([]int, 0, 26)
	for _, v := range m {
		count = append(count, v)
	}
	sort.Ints(count)
	result, countLen := -1, len(count)
	if count[countLen-1]-count[0] <= k {
		return 0
	}
	for i := 0; i < countLen; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += count[j]
		}
		for j := countLen - 1; j > i; j-- {
			if count[j]-count[i] <= k {
				break
			}
			sum += count[j] - count[i] - k
		}
		if result == -1 {
			result = sum
		} else {
			result = min(result, sum)
		}
	}
	return result
}

func main() {
	fmt.Println(minimumDeletions("aabcaba", 0))
	fmt.Println(minimumDeletions("dabdcbdcdcd", 2))
	fmt.Println(minimumDeletions("aaabaaa", 2))
	fmt.Println(minimumDeletions("itatwtiwwi", 1))
}
