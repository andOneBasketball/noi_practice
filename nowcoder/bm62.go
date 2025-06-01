package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	// write code here
	dp := make([]int, n+1)
	if n == 1 || n == 2 {
		return 1
	}
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibonacci(n))
}
