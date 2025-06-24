package main

import (
	"fmt"
)

const mod = 1e9 + 7

func comb(n, k int) int {
	if k > n/2 {
		k = n - k
	}
	x, y := 1, 1
	for i := 0; i < k; i++ {
		x = x * (n - i) % mod
		y = y * (i + 1) % mod
	}

	return x / y
}

func pow(n, k int) int {
	result := 1
	for i := 0; i < k; i++ {
		result = result * n % mod
	}
	return result
}

func countGoodArrays(n int, m int, k int) int {
	// C(n-1, k)
	fmt.Println(comb(n-1, k), m, pow(m-1, n-1-k))
	return comb(n-1, k) * m % mod * pow(m-1, n-1-k) % mod
}

func main() {
	fmt.Println(countGoodArrays(3, 2, 1))
	fmt.Println(countGoodArrays(4, 2, 2))
	fmt.Println(countGoodArrays(5, 2, 0))
	fmt.Println(countGoodArrays(5, 2, 2))
	fmt.Println(countGoodArrays(5581, 58624, 4766))
}
