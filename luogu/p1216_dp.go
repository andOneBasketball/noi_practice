package main

import (
	"fmt"
)

func calcMaxLen(arr [][]int, num int) int {
	i := num - 2
	for i >= 0 {
		lineLen := len(arr[i])
		for j := 0; j < lineLen; j++ {
			if arr[i][j]+arr[i+1][j] < arr[i][j]+arr[i+1][j+1] {
				arr[i][j] = arr[i][j] + arr[i+1][j+1]
			} else {
				arr[i][j] = arr[i][j] + arr[i+1][j]
			}
		}
		i--
	}
	return arr[0][0]
}

func main() {
	/*
		f, _ := os.Open("p1216.in")
		defer f.Close()
		os.Stdin = f
	*/

	num := 0
	fmt.Scan(&num)
	arr := make([][]int, num, num)
	for i := 0; i < num; i++ {
		arr[i] = make([]int, 0, num)
	}
	//fmt.Println(arr, num)
	line, data := 1, 0
	for i := 0; i < num; i++ {
		for j := 0; j < line; j++ {
			fmt.Scan(&data)
			arr[i] = append(arr[i], data)
		}
		line++
	}
	//fmt.Println(arr)
	fmt.Println(calcMaxLen(arr, num))
}
