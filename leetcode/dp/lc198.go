package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	numsLen := len(nums)
	maxData := -1
	dp := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if i > 2 {
			dp[i] = max(nums[i]+dp[i-2], nums[i]+dp[i-3])
		} else if i == 2 {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = nums[i]
		}
		maxData = max(maxData, dp[i])
	}
	return maxData
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
