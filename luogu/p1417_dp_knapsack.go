package main

import (
	"fmt"
	"sort"
)

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcMaxValue(abc [][3]int, t, n int) int {
	sort.Slice(abc, func(i, j int) bool {
		return func(i, j [3]int) bool {
			// 先i 后 j
			v1 := i[0] - i[1]*i[2] + j[0] - j[1]*(i[2]+j[2])
			// 先j 后 i
			v2 := j[0] - j[1]*j[2] + i[0] - i[1]*(i[2]+j[2])
			return v1 > v2
		}(abc[i], abc[j])
	})

	dp := make([]int, t+1)
	for i := 0; i < n; i++ {
		for j := t; j >= abc[i][2]; j-- {
			dp[j] = max(dp[j], dp[j-abc[i][2]]+abc[i][0]-abc[i][1]*j)
		}
	}
	v := 0
	for i := 0; i < t+1; i++ {
		v = maxValue(v, dp[i])
	}

	return v
}

func main() {
	/*
		f, _ := os.Open("p1417.in")
		defer f.Close()
		os.Stdin = f
	*/
	var t, n int
	fmt.Scan(&t, &n)
	aibici := make([][3]int, n)
	for i := 0; i < 3; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&aibici[j][i])
		}
	}
	//fmt.Println(aibici, t, n)
	fmt.Println(calcMaxValue(aibici, t, n))
}
