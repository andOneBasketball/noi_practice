package main

import (
	"fmt"
	"sort"
)

func maxProfit(n int, edges [][]int, score []int) int {
	nodeMap := make(map[int][]int, n)
	edgeMap := make(map[int]int, n)
	for i := 0; i < len(edges); i++ {
		v1, v2 := edges[i][0], edges[i][1]
		if _, ok := nodeMap[v1]; !ok {
			nodeMap[v1] = make([]int, 0, n)
		}
		nodeMap[v1] = append(nodeMap[v1], v2)
		edgeMap[v2]++
	}
	//fmt.Println(edgeMap)

	// 从度为0的节点开始
	result := 0
	list := make([]int, 0, n)
	s := make([]int, 0, n)
	flag := false
	for i := 0; i < n; i++ {
		if _, ok := edgeMap[i]; !ok {
			s = append(s, score[i])
			list = append(list, i+1)
			flag = true
			nodeList, nodeOk := nodeMap[i]
			if !nodeOk {
				continue
			}
			for _, v := range nodeList {
				if edgeMap[v] > 0 {
					edgeMap[v]--
				}
			}
		}
	}
	if flag {
		sort.Ints(list)
		sort.Ints(s)
		for i := 0; i < len(list); i++ {
			result += list[i] * s[i]
		}
		fmt.Println(list, s, result, edgeMap)
	}

	for {
		flag := false
		list = list[0:0]
		s = s[0:0]
		//fmt.Println(list, s, edgeMap)
		for k, v := range edgeMap {
			if v == 0 {
				s = append(s, score[k])
				//result += (k + 1) * score[k]
				list = append(list, k)
				flag = true
			}
		}
		sort.Ints(list)
		sort.Ints(s)
		for i := 0; i < len(list); i++ {
			//fmt.Println(list, s, result)
			result += (list[i] + 1) * s[i]
		}
		for _, v := range list {
			delete(edgeMap, v)
			nodeList, nodeOk := nodeMap[v]
			if !nodeOk {
				continue
			}
			for _, vv := range nodeList {
				if edgeMap[vv] > 0 {
					edgeMap[vv]--
				}
			}
		}
		fmt.Println(list, s, result, edgeMap)
		if !flag {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(maxProfit(2, [][]int{[]int{0, 1}}, []int{2, 3}))
	fmt.Println(maxProfit(3, [][]int{[]int{0, 1}, []int{0, 2}}, []int{1, 6, 3}))
	fmt.Println(maxProfit(3, [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 2}}, []int{61902, 56599, 31791}))
	fmt.Println(maxProfit(3, [][]int{[]int{0, 1}}, []int{95934, 94854, 7360}))
}
