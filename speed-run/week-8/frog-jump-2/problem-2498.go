package frogjump2

func maxJump(stones []int) int {
	ans := stones[1] - stones[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 2; i < len(stones); i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}

	return ans

}
