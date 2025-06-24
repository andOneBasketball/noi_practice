package main

import "fmt"

func findCoins(numWays []int) []int {
	dp := make([][3]int, 0, 101)
	for pos, v := range numWays {
		if v == 0 {
			continue
		} else if v == 1 {
			dp = append(dp, [3]int{pos + 1, 1, 1})
			continue
		}
		tmp := 0
		for i := 1; i*2 <= pos+1; i++ {
			if dp[i][0] >= 1 && dp[pos+1-i][0] >= 1 {
				tmp += dp[i][0]
			}
		}
		//fmt.Println(pos, v, tmp, dp)
		if tmp == v-1 {
			dp[pos+1] = [2]int{tmp + 1, 1}
		} else if tmp == v {
			dp[pos+1] = [2]int{tmp, 0}
		} else if tmp < v-1 {
			return []int{}
		}
	}

	result := make([]int, 0, 101)
	for i := 1; i < 101; i++ {
		if dp[i][0] >= 1 && dp[i][1] == 1 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println(findCoins([]int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}))
	fmt.Println(findCoins([]int{1, 2, 2, 3, 4}))
	fmt.Println(findCoins([]int{1, 2, 3, 4, 15}))
}
