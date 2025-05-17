package main

import (
	"fmt"
)

// 可以通过记录达标字符的数量来统计是否是覆盖子串
func tExist(chMap, tMap map[byte]int) bool {
	for tK, vK := range tMap {
		if data, ok := chMap[tK]; !ok || data < vK {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	result := ""
	minRes := ""
	chMap := make(map[byte]int, len(t))
	tMap := make(map[byte]int, len(t))
	for _, v := range t {
		tMap[byte(v)]++
	}
	existCount := len(tMap)
	currCount := 0
	calcChMap := make(map[byte]struct{}, len(t))
	i, j := 0, 0
	for j < len(s) {
		if _, ok := tMap[s[j]]; ok {
			chMap[s[j]]++
			_, chExist := calcChMap[s[j]]
			if chMap[s[j]] >= tMap[s[j]] && !chExist {
				currCount++
				calcChMap[s[j]] = struct{}{}
			}
			j++
			if j > len(s) {
				break
			}
			//fmt.Printf("j: %d, chMap: %v\n", j, chMap)
			if currCount == existCount {
				result = s[i:j]
				if minRes == "" || len(result) < len(minRes) {
					minRes = result
				}
				for i < j {
					//fmt.Printf("i: %d, j: %d, result: %s, s[j]: %c, s[i]: %c, chMap: %v, tMap: %v\n", i, j, result, s[j], s[i], chMap, tMap)
					if _, ok := chMap[s[i]]; ok {
						chMap[s[i]]--
						if chMap[s[i]] < tMap[s[i]] {
							delete(calcChMap, s[i])
							currCount--
						}
						if currCount != existCount {
							i++
							break
						}
					}
					i++
					result = s[i:j]
					if len(result) < len(minRes) {
						minRes = result
					}
				}
			}
			continue
		}
		j++
	}
	return minRes
}

func main() {
	/*
		fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
		fmt.Println(minWindow("a", "a"))
		fmt.Println(minWindow("a", "aa"))
	*/
	//fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
