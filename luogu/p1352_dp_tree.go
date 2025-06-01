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

func calcHappy(arr []int, tree map[int][]int, n int) int {
	stack := make([]int, 0, n+1)
	// 记录node的领导
	tree2 := make(map[int]int, 0)
	for node, v := range tree {
		if len(v) == 0 {
			stack = append(stack, node)
		} else {
			for _, idx := range v {
				tree2[idx] = node
			}
		}
	}
	//fmt.Println(arr, tree, tree2, stack)
	dp := make([][2]int, n+1)
	for len(stack) > 0 {
		newStack := make([]int, 0, n+1)
		flag := make(map[int]bool, 0)
		for _, v := range stack {
			dp[v][0] = 0
			dp[v][1] = arr[v]
			if len(tree[v]) > 0 {
				for _, i := range tree[v] {
					/*
						// 不知道 first code 时咋想的，为啥顺手就写成了 arr[dp[i][x]]
						dp[v][0] += arr[dp[i][1]]
						dp[v][1] += arr[dp[i][0]]

						// 快乐指数可能为负数，需要考虑来或者不来。
						dp[v][0] += dp[i][1]
						dp[v][1] += dp[i][0]

						// 但来的时候，下属只能不来
						dp[v][0] = maxValue(dp[v][0], dp[v][0]+dp[i][1])
						dp[v][1] = maxValue(dp[v][1], dp[v][1]+dp[i][0])
					*/
					dp[v][0] = maxValue(dp[v][0]+dp[i][0], dp[v][0]+dp[i][1])
					dp[v][1] += dp[i][0]
				}
			}

			if _, ok := tree2[v]; ok && !flag[tree2[v]] {
				newStack = append(newStack, tree2[v])
				flag[tree2[v]] = true
			}
		}
		//fmt.Println(stack, dp)
		stack = append([]int{}, newStack...)
		//fmt.Println(dp)
	}
	maxHappy := -1
	for _, v := range dp {
		maxHappy = maxValue(maxHappy, v[0])
		maxHappy = maxValue(maxHappy, v[1])
	}
	return maxHappy
}

func main() {
	/*
		f, _ := os.Open("p1352_4.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n int
	fmt.Scan(&n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}
	var l, k int

	tree := make(map[int][]int, n)
	for i := n + 1; i < 2*n; i++ {
		fmt.Scan(&l, &k)
		if _, ok := tree[k]; !ok {
			tree[k] = make([]int, 0, 1)
		}
		tree[k] = append(tree[k], l)
	}
	// 补充不是领导的节点
	for i := 1; i <= n; i++ {
		if _, ok := tree[i]; ok {
			continue
		}
		tree[i] = make([]int, 0)
	}
	fmt.Println(calcHappy(arr, tree, n))
}
