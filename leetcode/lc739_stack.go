package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	stack := make([][2]int, 0, len(temperatures))
	result := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(stack) > 0 && v > stack[len(stack)-1][0] {
			result[stack[len(stack)-1][1]] = i - stack[len(stack)-1][1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, [2]int{v, i})
	}
	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
