package main

import (
	"fmt"
)

func jump(nums []int) int {
	if nums[0] == 0 || len(nums) == 1 {
		return 0
	}
	count := 0
	i := 0
	for i < len(nums) {
		//fmt.Println(i)
		if nums[i] > 0 {
			// maxJump 每次需要重新初始化
			maxJump := [2]int{0, 0}
			count++

			if i+nums[i] >= len(nums)-1 {
				return count
			}
			// 边界值的确认
			for j := 1; j <= nums[i]; j++ {
				pos := j + i
				// 如果一样选更远的位置，切记是>=
				// 当前位置+能前往的最大位置
				if nums[pos]+pos >= maxJump[0] {
					maxJump[0] = nums[pos] + pos
					maxJump[1] = pos
				}
			}
			i = maxJump[1]
			//fmt.Printf("nums[%d]: %d, maxJump: %v\n", i, nums[i], maxJump)
		}
	}
	return count
}

/*
// leetcode 官方题解, 顺序轮询，每次确定边界，在边界内找到能跳跃到的最大位置，更新边界，每次更新边界时 step++
func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
*/

func jumpStd(nums []int) int {
	step := 0
	endPos := 0
	maxPos := 0
	// len(nums)-1 很关键，决定终点是否会触发，当i为终点时可能会触发 step++，而原本是不需要触发的
	for i, v := range nums[0 : len(nums)-1] {
		if i+v > maxPos {
			maxPos = i + v
		}
		if i == endPos {
			step++
			endPos = maxPos
		}
	}
	return step
}

func main() {
	/*
		fmt.Println(jump([]int{2, 3, 1, 1, 4}))
		fmt.Println(jump([]int{2, 3, 0, 1, 4}))
		fmt.Println(jump([]int{1, 2, 3}))
		fmt.Println(jump([]int{1, 2, 1, 1, 1}))
	*/
	//fmt.Println(jump([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}))

	//fmt.Println(jumpStd([]int{2, 3, 1, 1, 4}))
	fmt.Println(jumpStd([]int{1, 2, 3}))
}
