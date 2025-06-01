package main

import "fmt"

func pow(data, n int) int {
	result := data
	for i := 1; i < n; i++ {
		result *= data
	}
	return result
}

// arr: 低位在前,高位在后 eg: 576 []int{6, 7, 5}
func calcNum(arr []int, l int, num int) []int {
	result := make([]int, 10)
	dp := make([]int, l, l)
	dp[0] = 1
	for i := 0; i <= arr[0]; i++ {
		result[i] += 1
	}

	for i := 1; i < l; i++ {
		dp[i] = dp[i-1]*10 + pow(10, i)
		for j := 0; j < 10; j++ {
			result[j] += arr[i] * dp[i-1]
		}
		// eg: 567 0~4 执行以下操作
		for j := 0; j < arr[i]; j++ {
			result[j] += pow(10, i)
		}
		// 补充统筹临界位的数字
		mod := num % pow(10, i)
		result[arr[i]] += mod + 1

		// 处理前导0
		result[0] -= pow(10, i)
	}

	return result
}

func formatData(arr *[]int, data int) {
	if data == 0 {
		*arr = append(*arr, 0)
	}
	for data > 0 {
		mod := data % 10
		*arr = append(*arr, mod)
		data /= 10
	}
}

func calcDigtalNum(a, b int) []int {
	num := make([]int, 10)
	arrA, arrB := make([]int, 0, 15), make([]int, 0, 15)
	formatData(&arrA, a-1)
	formatData(&arrB, b)
	//fmt.Println(arrA, a-1)
	//fmt.Println(arrB, b)
	numA := calcNum(arrA, len(arrA), a-1)
	numB := calcNum(arrB, len(arrB), b)
	//fmt.Println(a-1, numA, b, numB)
	for i := 0; i < 10; i++ {
		num[i] = numB[i] - numA[i]
	}
	//fmt.Println(num, numA, numB)
	return num
}

func main() {
	var a, b int
	//a, b = 6, 976
	fmt.Scan(&a, &b)
	result := calcDigtalNum(a, b)
	for i, d := range result {
		if i == 9 {
			fmt.Println(d)
		} else {
			fmt.Printf("%d ", d)
		}
	}
}
