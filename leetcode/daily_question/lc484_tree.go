package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getStep(curr int, n int) int {
	left, right := curr, curr
	step := 0
	for left <= n {
		step += min(right, n) - left + 1
		left *= 10
		right = right*10 + 9
	}
	return step
}

func findKthNumber(n int, k int) int {
	curr := 1
	k--
	for k > 0 {
		step := getStep(curr, n)
		//fmt.Println(k, curr, n, step)
		if k >= step {
			curr++
			k -= step
		} else {
			k--
			curr *= 10
		}
	}
	return curr
}

func main() {
	fmt.Println(findKthNumber(13, 2))
}
