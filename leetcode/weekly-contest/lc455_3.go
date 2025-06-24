package main

import "fmt"

func minIncrease(n int, edges [][]int, cost []int) int {
	edgesLen := len(edges)
	m := make(map[int][]int, edgesLen)
	for _, v := range edges {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = make([]int, 0, 16)
		}
		m[v[0]] = append(m[v[0]], v[1])
	}
	maxValue := 0
}

func main() {
	fmt.Println(minIncrease(3, [][]int{[]int{0, 1}, []int{0, 2}}, []int{2, 1, 3}))
	fmt.Println(minIncrease(3, [][]int{[]int{0, 1}, []int{1, 2}}, []int{5, 1, 4}))
	fmt.Println(minIncrease(5, [][]int{[]int{0, 4}, []int{0, 1}, []int{1, 2}, []int{1, 3}}, []int{3, 4, 1, 1, 7}))
}
