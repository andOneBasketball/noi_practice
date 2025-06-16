package main

import (
	"fmt"
	"math"
)

func maxDifference(s string) int {
	sLen := len(s)
	m := make(map[byte]int, sLen)
	for i := 0; i < sLen; i++ {
		m[s[i]]++
	}
	max, min := -1, math.MaxInt
	for _, v := range m {
		if v%2 == 0 && v < min {
			min = v
		} else if v%2 == 1 && v > max {
			max = v
		}
	}
	return max - min
}

func main() {
	fmt.Println(maxDifference("aaaaabbc"))
}
