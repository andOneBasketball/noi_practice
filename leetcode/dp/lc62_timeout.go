package main

import "fmt"

func uniquePathsTimeout(m int, n int) int {
	step := [][3]int{
		[3]int{0, 1, 1},
		[3]int{1, 0, 2},
	}
	// 1 表示横坐标，2表示纵坐标
	stack := [][3]int{
		[3]int{0, 1, 1},
		[3]int{1, 0, 2},
	}
	count := 2
	if m == 1 || n == 1 {
		count = 1
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		for _, s := range step {
			x, y := v[0]+s[0], v[1]+s[1]
			if v[2] != s[2] && x < m && y < n {
				stack = append(stack, [3]int{x, y, s[2]})
				if (x == m-1 && s[2] == 1) || (y == n-1 && s[2] == 2) {
					continue
				}
				count++
			} else if v[2] == s[2] && x < m && y < n {
				stack = append(stack, [3]int{x, y, s[2]})
			}
		}
	}
	return count
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {

	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 3))
	fmt.Println(uniquePaths(1, 2))

	//fmt.Println(uniquePaths(23, 12))
}
