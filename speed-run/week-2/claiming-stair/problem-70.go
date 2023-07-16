package claimingstair

func climbStairs(n int) int {
	step1, step2 := 0, 1
	for i := 2; i <= n; i++ {
		step1, step2 = step2, step2+step1
	}

	return step2
}
