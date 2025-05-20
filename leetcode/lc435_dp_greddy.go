package main

import (
	"fmt"
	"sort"
)

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp 会 timeout
func eraseOverlapIntervalsDp(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return func(a, b []int) bool {
			if a[0] < b[0] {
				return true
			}
			if a[0] == b[0] {
				return a[1] < b[1]
			}
			return false
		}(intervals[i], intervals[j])
	})

	maxArea := 0
	intervalsLen := len(intervals)
	dp := make([]int, intervalsLen)
	for i, curr := range intervals {
		dp[i] = 1
		temp := 0
		for j, pre := range intervals[0:i] {
			if curr[0] >= pre[1] {
				temp = maxNum(dp[j], temp)
			}
		}
		dp[i] += temp
		maxArea = maxNum(maxArea, dp[i])
	}
	return intervalsLen - maxArea
}

// 贪心
func eraseOverlapIntervals(intervals [][]int) int {
	maxArea := 0
	sort.Slice(intervals, func(i, j int) bool {
		return func(a, b []int) bool {
			if a[1] < b[1] {
				return true
			}
			if a[1] == b[1] {
				return a[0] < b[0]
			}
			return false
		}(intervals[i], intervals[j])
	})
	//fmt.Println(intervals)
	i, j, intervalsLen := 0, 0, len(intervals)
	for j < intervalsLen {
		if maxArea == 0 {
			maxArea = 1
		}

		if j+1 < intervalsLen && intervals[i][1] <= intervals[j+1][0] {
			maxArea++
			j++
			i = j
			continue
		}
		j++
	}
	return intervalsLen - maxArea
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}}))
}
