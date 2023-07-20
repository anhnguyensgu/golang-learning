package longestsubstringwithoutreaptingcharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l := 0
	ans := 1

	mapChar := [128]int{}
	mapChar[s[0]]++

	for r := 1; r < len(s); r++ {
		if mapChar[s[r]] > 0 {
			for mapChar[s[l]] == mapChar[s[r]] && l < r {
				mapChar[s[l]]--
				l++
			}
		}
		mapChar[s[r]]++
		ans = max(ans, r-l+1)
	}

	return ans
}
