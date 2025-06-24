package main

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	result := 1
	i, j, numsLen := 0, 0, len(nums)
	for j < numsLen {
		if nums[j]-nums[i] > k {
			i = j
			result++
			continue
		}
		j++
	}
	return result
}

func main() {
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
	fmt.Println(partitionArray([]int{1, 2, 3}, 1))
	fmt.Println(partitionArray([]int{2, 2, 4, 5}, 0))
}
