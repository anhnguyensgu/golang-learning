package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairsWithArray(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func minCostClimbingStairs(cost []int) int {
	// dp := make([]int, len(cost))
	// dp[0] = cost[0]
	// dp[1] = cost[1]
	first := cost[0]
	second := cost[1]
	for i := 2; i < len(cost); i++ {
		temp := second
		second = min(first, second) + cost[i]
		first = temp
		// dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	// return min(dp[len(dp)-1], dp[len(dp)-2])
	return min(first, second)
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	ans := minCostClimbingStairs(cost)
	fmt.Println(ans)
}
