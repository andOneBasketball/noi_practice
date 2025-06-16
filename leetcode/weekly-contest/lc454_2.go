package main

import (
	"fmt"
)

func specialTriplets(nums []int) int {
	sum := 0
	numsLen := len(nums)
	numsMap := make(map[int][]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if _, ok := numsMap[nums[i]]; !ok {
			numsMap[nums[i]] = make([]int, 0, 256)
		}
		numsMap[nums[i]] = append(numsMap[nums[i]], i)
	}
	//fmt.Println(numsMap)
	for k, v := range numsMap {
		ikPos, ok := numsMap[k*2]
		if !ok {
			continue
		}
		ikPosLen := len(ikPos)
		if ikPosLen < 2 {
			continue
		}

		if k != 2*k {
			for _, j := range v {
				if ikPos[0] > j || ikPos[ikPosLen-1] < j {
					continue
				}
				l, mid, r := 0, (ikPosLen-1)/2, ikPosLen-1
				//fmt.Println(sum, l, mid, r, ikPos, k, j)
				for {
					//fmt.Println(l, mid, r, ikPos, k, j)
					if ikPos[mid] < j && ikPos[mid+1] > j {
						break
					}
					if ikPos[mid] > j {
						r = mid
					} else {
						l = mid
					}
					mid = (r-l)/2 + l
				}
				sum += (mid + 1) * (ikPosLen - mid - 1)
			}
		} else {
			vLen := len(v) - 1
			for i := 1; i <= vLen/2; i++ {
				if i == vLen-i {
					sum += i * (vLen - i)
				} else {
					sum += i * (vLen - i) * 2
				}
			}
		}
	}

	return sum % (1e9 + 7)
}

func main() {
	fmt.Println(specialTriplets([]int{84, 2, 93, 1, 2, 2, 26}))
	/*
		fmt.Println(specialTriplets([]int{6, 3, 6}))
		fmt.Println(specialTriplets([]int{0, 1, 0, 0}))
		fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4}))

		fmt.Println(specialTriplets([]int{0, 0, 0, 0, 0}))

		fmt.Println(specialTriplets([]int{0, 0, 0, 0, 0}))
	*/

}
