package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findKDistantIndices(nums []int, key int, k int) []int {
	numsLen := len(nums)
	pos := make([]int, 0, numsLen)
	for i, v := range nums {
		if v == key {
			pos = append(pos, i)
		}
	}
	posLen := len(pos)
	result := make([]int, 0, numsLen)
	i, j := 0, 0
	for i < numsLen {
		if abs(i-pos[j]) <= k {
			result = append(result, i)
		} else if j+1 < posLen && abs(i-pos[j+1]) <= k {
			result = append(result, i)
			j++
		}
		i++
	}
	return result
}

func main() {
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
}
