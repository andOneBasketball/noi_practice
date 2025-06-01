package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	numsLen := len(nums)
	numsMap := make(map[int]bool, numsLen)

	nums = append(nums, nums...)
	result := make([]int, numsLen)
	stack := make([][2]int, 0, numsLen)
	for i := 0; i < numsLen*2; i++ {
		//fmt.Println(i, stack)
		if len(stack) > 0 {
			for len(stack) > 0 && nums[i] > stack[len(stack)-1][0] {
				pos := stack[len(stack)-1][1] % numsLen
				result[pos] = nums[i]
				stack = stack[0 : len(stack)-1]
				numsMap[pos] = true
			}
		}
		if !numsMap[i%numsLen] {
			stack = append(stack, [2]int{nums[i], i % numsLen})
		}
	}
	if len(stack) > 0 {
		for _, v := range stack {
			result[v[1]] = -1
		}
	}
	return result
}

func main() {
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 2, 1}))
}
