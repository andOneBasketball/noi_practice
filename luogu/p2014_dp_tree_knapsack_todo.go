package main

import "fmt"

func calcMaxPoint(arr []int, tree map[int][]int, n int) int {
	stack := make([]int, 0, n)
	tree2 := make([][2]int, n)
	for k, v := range tree {
		if len(v) == 0 {
			stack = append(stack, k)
		} else {
			for _, idx := range v {
				// 0 表示该 node 尚未被选上
				tree2[idx] = [2]int{k, 0}
			}
		}
	}
	dp := make([]int, n)
	for len(stack) > 0 {
		for _, node := range stack {
			for j := n
		}
	}
}

func main() {
	var n, m, t, data int
	fmt.Scan(&n, &m)
	arr := make([]int, n+1)
	tree := make(map[int][]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&t, &data)
		if t == 0 {
			if _, ok := tree[i]; !ok {
				tree[i] = make([]int, 0)
			}
			arr[i] = data
		} else {
			if _, ok := tree[t]; !ok {
				tree[t] = make([]int, 0, 1)
			}
			tree[t] = append(tree[t], i)
		}
	}
	fmt.Println(calcMaxPoint(arr, tree, n+1))
}
