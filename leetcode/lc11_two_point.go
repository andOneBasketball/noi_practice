package main

import (
	"fmt"
)

func maxArea(height []int) int {
	result := 0
	/*
		//timeout
		for i := 0; i < len(height); i++ {
			for j := i + 1; j < len(height); j++ {
				h := height[i]
				if h > height[j] {
					h = height[j]
				}
				area := (j - i) * h
				if result < area {
					result = area
				}
			}
		}
	*/
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	i, j := 0, len(height)
	for i < j {
		left, right := height[i], height[j-1]
		h := left
		i++
		if h > right {
			h = right
			j--
			i--
		}
		area := (j - i) * h
		if result < area {
			result = area
		}
	}

	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
