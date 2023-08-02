package coinchange

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	dp := make([]int, amount+1)
	dp[0] = 0

	for curAmount := 1; curAmount <= amount; curAmount++ {
		curMin := int(10e5)
		for _, coin := range coins {
			if curAmount-coin >= 0 {
				curMin = min(dp[curAmount-coin]+1, curMin)
			}
		}
		dp[curAmount] = curMin
	}

	ans := dp[len(dp)-1]
	if ans == int(10e5) {
		return -1
	}

	return ans
}
