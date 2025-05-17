package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	/*
		// 暴力求解，timeout
		for i, v := range gas {
			if v < cost[i] || v == 0 {
				continue
			}
			flag := true
			c := 0
			for j := 0; j < len(gas); j++ {
				pos := (j + i) % len(gas)
				c += gas[pos] - cost[pos]
				if c < 0 {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
		}
	*/

	currentNum, totalNum, idx := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		currentNum += gas[i] - cost[i]
		totalNum += gas[i] - cost[i]
		// 经历过的地点都不可能了，只能寄希望于后续的加油站
		// 如果某个地点开始跑完了后续所有加油站，并且油量最后有剩，那肯定可以跑完全程。totalNum >= 0 表明前面的亏损会在以后慢慢补回来，则肯定能跑完
		// currentNum  < 0 场景，表明A => B 不可行，那 A+1/A+2/A+3 => B 都不可行(A => A+1/A+2/A+3 都是有盈余的)
		if currentNum < 0 {
			idx = (i + 1) % len(gas)
			currentNum = 0
		}
	}
	// 油不够，直接返回-1
	if totalNum < 0 {
		return -1
	}
	return idx
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
