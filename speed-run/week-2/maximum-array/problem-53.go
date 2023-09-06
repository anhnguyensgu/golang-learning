package maximumarray

func maxSubArray(nums []int) int {
	ans := nums[0]
	with := ans

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums); i++ {
		with = max(with+nums[i], nums[i])
		ans = max(with, ans)
	}
	return ans
}
