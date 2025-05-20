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

func sliceLess(a, b []int) bool {
	minLen := len(a)
	if minLen > len(b) {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	if len(a) < len(b) {
		return true
	}
	return false
}

func getMinSolution(arr [][2]int, n, v int) []int {
	solutions := make([][]int, v+1)
	// 初始化切片时初始化的长度终止条件，为啥一开始写成了 i < n 。。。
	for i := 0; i < v+1; i++ {
		solutions[i] = make([]int, 0, v+1)
	}
	dp := make([]int, v+1)
	for i := 0; i < n; i++ {
		for j := v; j >= arr[i][0]; j-- {
			pickValue := dp[j-arr[i][0]] + arr[i][1]
			oldValue := dp[j]
			dp[j] = maxValue(dp[j], pickValue)
			if oldValue < pickValue {
				// 深拷贝方式
				newSolution := append([]int{}, solutions[j-arr[i][0]]...)
				newSolution = append(newSolution, i+1)
				solutions[j] = newSolution
				// 切记先复制方案, 浅拷贝存在问题，切片的底层数据空间是一致的，会引发数据赋值紊乱
				//solutions[j] = solutions[j-arr[i][0]]
				//copy(solutions[j], solutions[j-arr[i][0]])
				//solutions[j] = append(solutions[j], i+1)
			} else if oldValue == pickValue {
				newSolution := append([]int{}, solutions[j-arr[i][0]]...)
				newSolution = append(newSolution, i+1)
				//fmt.Printf("newSolution:%v, j:%d, arr[i][0]:%d, solutions:%v, i:%d\n", newSolution, j, arr[i][0], solutions, i)
				if !sliceLess(solutions[j], newSolution) {
					solutions[j] = newSolution
				}
			}
			//fmt.Printf("i: %d, j: %d, oldValue: %d, pickValue: %d, solutions[%d]=%v, arr[%d][0]=%d, solutions=%v\n", i, j, oldValue, pickValue, j, solutions[j], i, arr[i][0], solutions)
		}
	}
	//fmt.Println(dp[v])
	return solutions[v]
}

func main() {
	/*
		f, _ := os.Open("p12_2.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n, v int
	fmt.Scan(&n, &v)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	result := getMinSolution(arr, n, v)
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			fmt.Printf("%d", result[i])
		} else {
			fmt.Printf("%d ", result[i])
		}
	}
}
