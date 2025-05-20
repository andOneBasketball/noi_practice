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

func calcSolutionNum(arr [][2]int, n, v int) int {
	dp := make([]int, v+1)
	dpCnt := make([]int, v+1)
	for i := 0; i < v+1; i++ {
		dpCnt[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := v; j >= arr[i][0]; j-- {
			oldValue := dp[j]
			pickValue := dp[j-arr[i][0]] + arr[i][1]
			dp[j] = maxValue(dp[j], pickValue)
			cnt := 0
			//fmt.Printf("arr[%d]=%v, dpCnt[%d]=%d, dpCnt[%d]=%d\n", i, arr[i], j, dpCnt[j], j-arr[i][0], dpCnt[j-arr[i][0]])
			if oldValue == dp[j] && dp[j] != pickValue {
				cnt = dpCnt[j]
			} else if oldValue != dp[j] && dp[j] == pickValue {
				cnt = dpCnt[j-arr[i][0]]
			} else {
				cnt = dpCnt[j] + dpCnt[j-arr[i][0]]
			}
			dpCnt[j] = cnt % (1e9 + 7)
			//fmt.Printf("dpCnt[%d]=%d, dp[%d]=%d, oldValue: %d, cnt:%d, pickValue:%d\n", j, dpCnt[j], j, dp[j], oldValue, cnt, pickValue)
		}
	}
	//fmt.Println(dp[v], dpCnt[v])
	return dpCnt[v]
}

func main() {
	/*
		f, _ := os.Open("p11_3.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n, v int
	fmt.Scan(&n, &v)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	fmt.Println(calcSolutionNum(arr, n, v))
}
