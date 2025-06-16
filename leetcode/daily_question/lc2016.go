package main

import "fmt"

func maximumDifference(nums []int) int {
	result := -1
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			odd := nums[j] - nums[i]
			if odd > result && odd > 0 {
				result = odd
			}
		}
	}
	return result
}

func main() {
	fmt.Println(maximumDifference([]int{7, 1, 5, 4}))
	fmt.Println(maximumDifference([]int{9, 4, 3, 2}))
	fmt.Println(maximumDifference([]int{48, 45, 26, 26, 23, 15}))
}
