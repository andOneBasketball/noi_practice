package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxDistance(s string, k int) int {
	cntN, cntS, cntW, cntE, ans := 0, 0, 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'N' {
			cntN++
		} else if s[i] == 'S' {
			cntS++
		} else if s[i] == 'W' {
			cntW++
		} else {
			cntE++
		}
		minY, minX := min(cntN, cntS), min(cntW, cntE)
		oddY, oddX := abs(cntN-cntS), abs(cntW-cntE)
		if k > minY {
			oddY += 2 * minY
			left := k - minY
			if left > minX {
				oddX += 2 * minX
			} else {
				oddX += 2 * left
			}
		} else {
			oddY += 2 * k
		}
		ans = max(ans, oddY+oddX)
	}
	return ans
}

func main() {
	fmt.Println(maxDistance("NWSE", 1))
	fmt.Println(maxDistance("NSWWEW", 3))
	// 以下场景考虑缺失
	fmt.Println(maxDistance("WSNS", 1))
	fmt.Println(maxDistance("WSNENN", 2))
	fmt.Println(maxDistance("NNSSEW", 3))
	fmt.Println(maxDistance("NEWSEW", 2))
}
