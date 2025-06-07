package main

import (
	"fmt"
)

func calc(arr []int, u, v, p int, mData map[[3]int]int) int {
	if data, ok := mData[[3]int{u, v, p}]; ok {
		return data
	}
	maxValue := 0
	for i := u; i <= v; i++ {
		mod := arr[i] % p
		if mod > maxValue {
			maxValue = mod
		}
	}
	mData[[3]int{u, v, p}] = maxValue
	return maxValue
}

func main() {
	/*
		f, _ := os.Open("p6141.in")
		defer f.Close()
		os.Stdin = f
	*/
	var n, m, u, v, p int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	result := make([]int, 0, m)
	mData := make(map[[3]int]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v, &p)
		result = append(result, calc(arr, u, v, p, mData))
	}
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			fmt.Printf("%d", result[i])
		} else {
			fmt.Println(result[i])
		}
	}
}
