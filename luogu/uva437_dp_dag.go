package main

import (
	"fmt"
)

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func babyLonSub(arr [][3]int, idx int, solutionNum int, dp [][3]int, n int) int {
	if dp[idx][solutionNum] != -1 {
		return dp[idx][solutionNum]
	}
	dp[idx][solutionNum] = 0

	length, width := -1, -1
	if solutionNum == 0 {
		length = arr[idx][0]
		width = arr[idx][1]
	} else if solutionNum == 1 {
		length = arr[idx][0]
		width = arr[idx][2]
	} else if solutionNum == 2 {
		length = arr[idx][1]
		width = arr[idx][2]
	}

	for i := 0; i < n; i++ {
		// 递归求 dp[idx][solutionNum]
		if (arr[i][0] < length && arr[i][1] < width) || (arr[i][1] < length && arr[i][0] < width) {
			dp[idx][solutionNum] = maxValue(dp[idx][solutionNum], babyLonSub(arr, i, 0, dp, n)+arr[i][2])
		}
		if (arr[i][0] < length && arr[i][2] < width) || (arr[i][2] < length && arr[i][0] < width) {
			dp[idx][solutionNum] = maxValue(dp[idx][solutionNum], babyLonSub(arr, i, 1, dp, n)+arr[i][1])
		}
		if (arr[i][1] < length && arr[i][2] < width) || (arr[i][2] < length && arr[i][1] < width) {
			dp[idx][solutionNum] = maxValue(dp[idx][solutionNum], babyLonSub(arr, i, 2, dp, n)+arr[i][0])
		}
	}
	return dp[idx][solutionNum]
}

func babyLon(arr [][3]int, n int) int {
	// dp 需要赋初值，然后记忆搜索直接返回
	dp := make([][3]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = -1
		dp[i][1] = -1
		dp[i][2] = -1
	}

	height := 0
	for i := 0; i < n; i++ {
		height = maxValue(height, babyLonSub(arr, i, 0, dp, n)+arr[i][2])
		height = maxValue(height, babyLonSub(arr, i, 1, dp, n)+arr[i][1])
		height = maxValue(height, babyLonSub(arr, i, 2, dp, n)+arr[i][0])
	}
	return height
}

func main() {
	/*
		f, _ := os.Open("uva437.in")
		defer f.Close()
		os.Stdin = f
	*/
	idx := 1
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		arr := make([][3]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i][0], &arr[i][1], &arr[i][2])
		}
		height := babyLon(arr, n)
		fmt.Printf("Case %d: maximum height = %d\n", idx, height)

		idx++
	}
}
