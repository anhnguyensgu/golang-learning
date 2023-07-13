package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	curMin, curMax := prices[0], prices[0]
	ans := 0

	for _, p := range prices {
		if curMin > p {
			ans = max(ans, curMax-curMin)
			curMax = p
			curMin = p
		} else if curMax < p {
			curMax = p
		}
	}

	return max(ans, curMax-curMin)
}
