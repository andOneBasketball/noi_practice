package main

import (
	"fmt"
	"sort"
)

func lexicalOrder(n int) []int {
	result := make([]int, 0, n)
	num := 1
	for i := 1; i <= n; i++ {
		result = append(result, num)
		if num*10 <= n {
			num *= 10
		} else {
			// 注意是 for 而不是 if
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num += 1
		}
	}

	return result
}

func main() {
	result := lexicalOrder(192)
	fmt.Println(result, len(result))
	sort.Ints(result)
	fmt.Println(result)
}
