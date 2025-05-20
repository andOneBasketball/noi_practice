package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcIntervalTime(beginTime, endTime string) int {
	beginArr := strings.Split(beginTime, ":")
	endArr := strings.Split(endTime, ":")
	beginH, _ := strconv.Atoi(beginArr[0])
	beginM, _ := strconv.Atoi(beginArr[1])
	endH, _ := strconv.Atoi(endArr[0])
	endM, _ := strconv.Atoi(endArr[1])
	return endH*60 + endM - beginH*60 - beginM
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDataC(tree [][3]int, intervalTime int, num int) int {
	dp := make([]int, intervalTime+1)
	for i := 0; i < num; i++ {
		if tree[i][2] == 0 {
			for t := tree[i][0]; t <= intervalTime; t++ {
				dp[t] = maxValue(dp[t], dp[t-tree[i][0]]+tree[i][1])
			}
		} else if tree[i][2] == 1 {
			for t := intervalTime; t >= tree[i][0]; t-- {
				dp[t] = maxValue(dp[t], dp[t-tree[i][0]]+tree[i][1])
			}
		} else {
			for t := intervalTime; t >= tree[i][0]; t-- {
				for k := 1; k <= tree[i][2] && k*tree[i][0] <= t; k++ {
					dp[t] = maxValue(dp[t], dp[t-k*tree[i][0]]+k*tree[i][1])
				}
			}
		}
	}
	return dp[intervalTime]
}

func main() {
	/*
		f, _ := os.Open("p1833.in")
		defer f.Close()
		os.Stdin = f
	*/
	var beginTime, endTime string
	var n int
	fmt.Scan(&beginTime, &endTime, &n)
	intervalTime := calcIntervalTime(beginTime, endTime)
	tree := make([][3]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tree[i][0], &tree[i][1], &tree[i][2])
	}
	fmt.Println(maxDataC(tree, intervalTime, n))
}
