package main

import (
	"fmt"
	"os"
)

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clacMaxValue(arr [][][2]int, n int, weight int) int {
	dp := make([]int, weight+1)
	for i := 0; i < n+1; i++ {
		if len(arr[i]) == 0 {
			continue
		}
		for j := weight; j >= 0; j-- {
			// 这一轮循环很关键，组内物品循环为最后一轮，保障每个分组只能取一个物品
			for k := 0; k < len(arr[i]); k++ {
				if j >= arr[i][k][0] {
					dp[j] = maxNum(dp[j], dp[j-arr[i][k][0]]+arr[i][k][1])
				}
			}
		}
	}

	return dp[weight]
}

func main() {

	f, _ := os.Open("p1757.in")
	defer f.Close()
	os.Stdin = f

	var n, weight, a, b, c int
	fmt.Scan(&weight, &n)
	arr := make([][][2]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = make([][2]int, 0, n)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b, &c)
		//fmt.Println(a, b, c)
		arr[c] = append(arr[c], [2]int{a, b})
	}

	fmt.Println(clacMaxValue(arr, n, weight))
}
