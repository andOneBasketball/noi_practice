package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	numsMap := make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		numsMap[nums[i]]++
	}

	result := make([][]int, 0, numsLen)
	m := make(map[[3]int]struct{}, numsLen)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			data := 0 - nums[i] - nums[j]
			//fmt.Println(data, numsMap)
			if data == nums[i] && data == nums[j] && numsMap[data] < 3 {
				continue
			} else if data == nums[i] && data != nums[j] && numsMap[data] < 2 {
				continue
			} else if data != nums[i] && data == nums[j] && numsMap[data] < 2 {
				continue
			} else if data != nums[i] && data != nums[j] && numsMap[data] < 1 {
				continue
			} else {
				dest := []int{nums[i], nums[j], data}
				sort.Ints(dest)
				if _, ok := m[[3]int{dest[0], dest[1], dest[2]}]; !ok {
					result = append(result, []int{nums[i], nums[j], data})
					m[[3]int{dest[0], dest[1], dest[2]}] = struct{}{}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
