package main

import "fmt"

func prime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeFrequency(nums []int) bool {
	numsLen := len(nums)
	m := make(map[int]int, numsLen)
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		if prime(v) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkPrimeFrequency([]int{1, 2, 3, 4, 5, 4}))
	fmt.Println(checkPrimeFrequency([]int{1, 2, 3, 4, 5}))
	fmt.Println(checkPrimeFrequency([]int{2, 2, 2, 4, 4}))
}
