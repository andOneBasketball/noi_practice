package main

import (
	"fmt"
)

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcMaxNum(arr [][2]int, n, m, t int) int {
	dp := make([][]int, m+1)
	for i := 0; i < t+1; i++ {
		dp[i] = make([]int, t+1)
	}
	//fmt.Println(dp)
	for x := 0; x < n; x++ {
		for y := m; y >= arr[x][0]; y-- {
			for z := t; z >= arr[x][1]; z-- {
				//fmt.Println(y, z, arr[x][0], arr[x][1])
				dp[y][z] = maxNum(dp[y][z], dp[y-arr[x][0]][z-arr[x][1]]+1)
			}
		}
	}
	return dp[m][t]
}

func main() {
	/*
		f, _ := os.Open("p1855.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n, m, t int
	fmt.Scan(&n, &m, &t)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	//fmt.Println(arr, n, m, t)
	fmt.Println(calcMaxNum(arr, n, m, t))
}
