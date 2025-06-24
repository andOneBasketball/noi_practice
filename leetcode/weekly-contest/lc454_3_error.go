package main

import (
	"fmt"
	"math"
	"sort"
)

func getMaX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int, m int) int {
	numsLen := len(nums)
	newNums := make([][2]int, numsLen)
	for i := 0; i < numsLen; i++ {
		newNums[i] = [2]int{nums[i], i}
	}
	sort.Slice(newNums, func(i, j int) bool {
		if newNums[i][0] == newNums[j][0] {
			return newNums[i][1] < newNums[j][1]
		}
		return newNums[i][0] < newNums[j][0]
	})
	sum := -math.MaxInt
	i, j := 0, numsLen-1
	if m == 1 {
		sum = getMaX(newNums[0][0]*newNums[0][0], newNums[numsLen-1][0]*newNums[numsLen-1][0])
	} else {
		for 
	}
	return sum
}

func main() {
	// Test cases can be added here to validate the function
	fmt.Println(maximumProduct([]int{-1, -9, 2, 3, -2, -3, 1}, 1))
	fmt.Println(maximumProduct([]int{1, 3, -5, 5, 6, -4}, 3))
	fmt.Println(maximumProduct([]int{2, -1, 2, -6, 5, 2, -5, 7}, 2))
}
