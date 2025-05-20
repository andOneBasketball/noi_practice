/*
在一个圆形操场的四周摆放 N 堆石子，现要将石子有次序地合并成一堆，规定每次只能选相邻的 2 堆合并成新的一堆，并将新的一堆的石子数，记为该次合并的得分。
试设计出一个算法,计算出将 N 堆石子合并成 1 堆的最小得分和最大得分。

输入格式
数据的第 1 行是正整数 N，表示有 N 堆石子。
第 2 行有 N 个整数，第 i 个整数 ai 表示第 i 堆石子的个数。

输出格式
输出共 2 行，第 1 行为最小得分，第 2 行为最大得分。
*/

package main

import (
	"fmt"
	"math"
)

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcValue(arr []int, n int) (int, int) {
	dpMax := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dpMax[i] = make([]int, 2*n)
	}
	dpMin := make([][][2]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dpMin[i] = make([][2]int, 2*n)
	}

	/*
		// 必须从尾部开始递推，从首部开始会导致过程中存在部分区间还未被计算
		for i := 0; i < 2*n-1; i++ {
			for j := i + 1; j < 2*n; j++ {
				for k := i; k < j; k++ {
	*/
	// 从尾开始遍历
	for i := 2*n - 2; i >= 0; i-- {
		for j := i + 1; j < 2*n; j++ {
			for k := i; k < j; k++ {
				//dpMax[i][j] = maxValue(dpMax[i][j], dpMax[i][k]+dpMax[k+1][j]+arr[j]-arr[i])
				//dpMin[i][j] = minValue(dpMin[i][j], dpMin[i][k]+dpMin[k+1][j]+arr[j]-arr[i])
				if i > 0 {
					dpMax[i][j] = maxValue(dpMax[i][j], dpMax[i][k]+dpMax[k+1][j]+arr[j]-arr[i-1])

					pickValue := dpMin[i][k][0] + dpMin[k+1][j][0] + arr[j] - arr[i-1]
					if dpMin[i][j][1] == 0 {
						dpMin[i][j][0] = pickValue
						dpMin[i][j][1] = 1
					} else {
						dpMin[i][j][0] = minValue(dpMin[i][j][0], pickValue)
					}
				} else {
					dpMax[i][j] = maxValue(dpMax[i][j], dpMax[i][k]+dpMax[k+1][j]+arr[j])

					pickValue := dpMin[i][k][0] + dpMin[k+1][j][0] + arr[j]
					if dpMin[i][j][1] == 0 {
						dpMin[i][j][0] = pickValue
					} else {
						dpMin[i][j][0] = minValue(dpMin[i][j][0], pickValue)
					}
				}
			}
		}
	}
	maxData, minData := -1, math.MaxInt32
	//fmt.Println(maxData, minData)
	for i := 0; i < n; i++ {
		maxData = maxValue(maxData, dpMax[i][i+n-1])
		minData = minValue(minData, dpMin[i][i+n-1][0])
	}
	return maxData, minData
}

func main() {
	/*
		f, _ := os.Open("p1880.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	zeroFlag := true
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] != 0 {
			zeroFlag = false
		}
	}
	//fmt.Println(arr)
	if !zeroFlag {
		arr = append(arr, arr...)
		for i := 1; i < 2*n; i++ {
			arr[i] += arr[i-1]
		}

		max, min := calcValue(arr, n)
		fmt.Println(min)
		fmt.Println(max)
	} else {
		fmt.Println(0)
		fmt.Println(0)
	}
}
