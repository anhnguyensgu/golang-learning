package containerwithmostwater

func maxArea(height []int) int {
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l, r := 0, len(height)

	for l < r {
		cur := (r - l) * min(height[l], height[r])
		ans = max(cur, ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
