package main

import "fmt"

type NumArray struct {
	ArrSum []int
}

func Constructor(nums []int) NumArray {
	numsLen := len(nums)
	arrSum := make([]int, numsLen)
	sum := 0
	for i := 0; i < numsLen; i++ {
		sum += nums[i]
		arrSum[i] = sum
	}

	return NumArray{
		ArrSum: arrSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.ArrSum[right]
	}
	return this.ArrSum[right] - this.ArrSum[left-1]
}

func main() {
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(n.SumRange(0, 2))
	fmt.Println(n.SumRange(2, 5))
	fmt.Println(n.SumRange(0, 5))
}
