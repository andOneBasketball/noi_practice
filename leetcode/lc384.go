package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	numsLen := len(this.nums)
	newNums := make([]int, numsLen)
	copy(newNums, this.nums)
	for i := 0; i < numsLen-1; i++ {
		pos := rand.Intn(numsLen-i) + i
		newNums[i], newNums[pos] = newNums[pos], newNums[i]
	}
	return newNums
}

func main() {
	s := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(s)
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}
