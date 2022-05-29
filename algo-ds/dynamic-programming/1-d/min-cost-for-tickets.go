package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func forward(days []int, costs []int) int {
	dp := make([]int, len(days))
	durations := []int{1, 7, 30}

	for i, day := range days {
		ans := 400000
		for j, duration := range durations {
			cur := i
			for cur >= 0 && days[cur]+duration > day {
				cur--
			}
			if cur < 0 {
				ans = min(costs[j], ans)
			} else {
				ans = min(dp[cur]+costs[j], ans)
			}
		}
		dp[i] = ans
	}
	return dp[len(days)-1]
}

func minCostTicketsWindowApproach(days []int, costs []int) int {
	dp := make([]int, len(days))
	durations := []int{1, 7, 30}

	for i := len(days) - 1; i >= 0; i-- {
		ans := 400000
		for j, duration := range durations {
			cur := i
			for cur < len(days) && days[cur] < days[i]+duration {
				cur++
			}
			if cur >= len(days) {
				ans = min(ans, costs[j])
			} else {
				ans = min(ans, dp[cur]+costs[j])
			}
		}
		dp[i] = ans
	}
	return dp[0]
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 400)
	plans := []int{1, 7, 30}
	daySet := map[int]bool{}
	for _, day := range days {
		daySet[day] = true
	}

	for i := 365; i >= 0; i-- {
		ans := 400000
		if daySet[i] {
			for j, plan := range plans {
				if i+plan <= 365 {
					ans = min(dp[i+plan]+costs[j], ans)
				} else {
					ans = min(costs[j], ans)
				}
			}
		} else {
			ans = dp[i+1]
		}
		if ans != 400000 {
			dp[i] = ans
		} else {
			dp[i] = 0
		}
	}
	return dp[1]
}

func main() {
	days := []int{1, 4, 6, 7, 8, 365}
	costs := []int{2, 7, 15}
	ans := forward(days, costs)
	fmt.Println(ans)
}
