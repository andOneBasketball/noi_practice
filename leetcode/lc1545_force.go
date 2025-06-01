package main

import (
	"fmt"
)

/*
0                1
011              1+1+1
0111001          3+1+3
01110011
1000110
0110001
011100110110001
0111001101100011011100100110001
*/

func main() {
	dp := make([]string, 30)
	dp[0] = "0"
	fmt.Println(dp[0])
	for i := 1; i < 10; i++ {

	}
}
