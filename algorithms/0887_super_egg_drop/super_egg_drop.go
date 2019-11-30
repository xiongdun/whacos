package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 5))
}

func superEggDrop(K int, N int) int {

	i, j := K+1, N+1
	dp := [i][j]int{}

	for m := 1; m <= N; m++ {

		dp[0][m] = 0

		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
			if dp[k][m] >= N {
				return m
			}
		}
	}

	return N
}
