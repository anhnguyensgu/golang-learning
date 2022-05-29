package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 2_147_483_648
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == 2_147_483_648 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1}
	amount := 0
	fmt.Printf("ans %d", coinChange(coins, amount))
}
