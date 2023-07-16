package longestpalindrome

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}
	charMap := [100]int{}
	for _, v := range s {
		charMap[v-'A']++
	}

	odd := 0
	even := 0
	for _, v := range charMap {
		if v%2 == 0 {
			even += v
		} else {
			odd++
			even += v - 1
		}
	}

	if odd > 0 {
		return even + 1
	}
	return even
}
