/*
第一行有两个整数，分别表示总钱数 n 和希望购买的物品个数 m。第 2 到第 (m+1) 行，每行三个整数，第 (i+1) 行的整数 vi，pi​，qi 分别表示第 i 件物品的价格、重要度以及它对应的的主件。如果 qi =0，表示该物品本身是主件。
*/

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

func treeKnapsack(arr [][][2]int, arrLen int, n int) int {
	dp := make([]int, n+1)
	for x := 0; x < arrLen; x++ {
		if len(arr[x]) == 0 || arr[x][0][0] == -1 {
			continue
		}

		for y := n; y >= 0; y-- {
			for z := 0; z < len(arr[x]); z++ {
				if y >= arr[x][z][0] {
					dp[y] = maxValue(dp[y], dp[y-arr[x][z][0]]+arr[x][z][1])
				}
			}
		}
	}
	return dp[n]
}

func formatKnapsack(arr [][][2]int, arrLen int) {
	for i := 0; i < arrLen; i++ {
		if len(arr[i]) == 0 || len(arr[i]) == 1 || arr[i][0][0] == -1 {
			continue
		}
		vList := make([][2]int, 0, len(arr[i])*2)
		for x := 1; x < len(arr[i]); x++ {
			price := arr[i][x][0]
			value := arr[i][x][1]
			tmpV := make([][2]int, 0, len(vList))
			for _, v := range vList {
				tmpV = append(tmpV, [2]int{price + v[0], value + v[1]})
			}
			vList = append(vList, tmpV...)
			vList = append(vList, [2]int{price, value})
			//fmt.Println(price, value, vList, len(arr[i]), i)
		}
		pos := 1
		for _, v := range vList {
			if pos < len(arr[i]) {
				arr[i][pos] = [2]int{v[0] + arr[i][0][0], v[1] + arr[i][0][1]}
				pos++
			} else {
				arr[i] = append(arr[i], [2]int{v[0] + arr[i][0][0], v[1] + arr[i][0][1]})
			}
		}
	}
}

func main() {
	/*
		f, _ := os.Open("p1064_2.in")
		defer f.Close()
		os.Stdin = f
	*/
	var p, n, a, b, c int
	fmt.Scan(&p, &n)
	arr := make([][][2]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = make([][2]int, 1, n+1)
		arr[i][0] = [2]int{-1, -1}
	}
	// group := 1 理解错了题目意思，分组就是直接递增分组
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b, &c)
		if c != 0 {
			arr[c] = append(arr[c], [2]int{a, a * b})
			continue
		}
		arr[i+1][0] = [2]int{a, a * b}
		//group++
	}
	//fmt.Println(arr)
	formatKnapsack(arr, n+1)
	//fmt.Println(arr)
	fmt.Println(treeKnapsack(arr, n+1, p))
}
