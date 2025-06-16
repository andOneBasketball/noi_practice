package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	numsLen := len(nums)
	nums1 := make([]int, 0, numsLen)
	nums2 := make([]int, 0, numsLen)
	numsMap := make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		numsMap[nums[i]]++
	}
	for k, v := range numsMap {
		if v > 1 {
			maxPos := v
			if v%2 == 1 {
				nums2 = append(nums2, k)
				maxPos--
			}
			for i := 0; i < maxPos; i++ {
				nums1 = append(nums1, k)
			}
		} else {
			nums2 = append(nums2, k)
		}
	}
	sort.Ints(nums2)
	nums1 = append(nums1, nums2...)
	if p == 0 {
		return 0
	}
	//fmt.Println(nums1)
	return nums1[p*2-1] - nums1[p*2-2]
}

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))
}
