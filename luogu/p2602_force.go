package main

import (
	"fmt"
)

func calc(num int) int {
	count := 0
	for i := 0; i <= num; i++ {
		j := i
		for {
			mod := j % 10
			if mod == 1 {
				count++
				//fmt.Printf("%d ", i)
			}
			j = j / 10
			if j == 0 {
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(calc(976))
}
