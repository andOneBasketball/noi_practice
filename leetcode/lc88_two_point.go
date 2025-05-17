package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, len(nums1))
	i, j := 0, 0
	for i < m || j < n {
		if len(nums1) > i && len(nums2) > j && i < m && nums1[i] <= nums2[j] {
			result[i+j] = nums1[i]
			i++
		}
		if len(nums1) > i && len(nums2) > j && j < n && nums1[i] > nums2[j] {
			result[i+j] = nums2[j]
			j++
		}
		// 继续处理未循环完的i 和 j,否则死循环
		if i == m && j < n {
			for j < n {
				result[i+j] = nums2[j]
				j++
			}
		}
		if j == n && i < m {
			for i < m {
				result[i+j] = nums1[i]
				i++
			}
		}
	}
	copy(nums1, result)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
