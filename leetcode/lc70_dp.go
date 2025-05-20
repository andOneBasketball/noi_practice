package main

import "fmt"

// 记忆搜索
func climbStairs(n int) int {
	stairsMap := make(map[int]int, 50)
	stairsMap[1] = 1
	stairsMap[2] = 2
	for i := 3; i <= n; i++ {
		stairsMap[i] = stairsMap[i-1] + stairsMap[i-2]
	}
	return stairsMap[n]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
