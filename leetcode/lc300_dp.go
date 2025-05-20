package main

import (
	"fmt"
)

func getPos(nums []int, curr int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	left, right := 0, numsLen
	if right-1 >= 0 && nums[right-1] < curr {
		return numsLen
	}
	for left < right {
		mid := left + (right-left)/2
		if curr > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 贪心算法，新增单调数组，单调数组的长度就是最长子序列的长度。
// 通过二分的方式寻找插入位置
func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	newNums := make([]int, 0, numsLen)

	for _, curr := range nums {
		pos := getPos(newNums, curr)
		if len(newNums) == 0 || pos == len(newNums) {
			newNums = append(newNums, curr)
		} else {
			newNums[pos] = curr
		}
		//fmt.Println(curr, pos, newNums)
	}
	return len(newNums)
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 时间复杂度 N^2
func lengthOfLISDp(nums []int) int {
	maxLen, numsLen := 0, len(nums)
	dp := make([]int, numsLen)
	for i, curr := range nums {
		dp[i] = 1
		temp := 0
		for j, pre := range nums[0:i] {
			if curr > pre {
				temp = maxNum(temp, dp[j])
			}
		}
		dp[i] += temp
		maxLen = maxNum(maxLen, dp[i])
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7}))
}
