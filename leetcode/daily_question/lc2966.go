package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	x, y, z, numsLen := 0, 1, 2, len(nums)
	result := make([][]int, 0, numsLen/3)
	for z < numsLen {
		if abs(nums[x]-nums[y]) <= k && abs(nums[x]-nums[z]) <= k && abs(nums[y]-nums[z]) <= k {
			result = append(result, []int{nums[x], nums[y], nums[z]})
		} else {
			return [][]int{}
		}
		x, y, z = x+3, y+3, z+3
	}
	return result
}

func main() {
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray([]int{2, 4, 2, 2, 5, 2}, 2))
}
