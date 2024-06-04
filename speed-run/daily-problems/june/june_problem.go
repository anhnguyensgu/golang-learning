package june

func appendCharacters(s string, t string) int {
	idx, subseqIdx := 0, 0
	for idx < len(s) && subseqIdx < len(t) {
		if s[idx] == t[subseqIdx] {
			subseqIdx++
		}
		idx++
	}
	if subseqIdx == len(t) {
		return 0
	}

	return len(t) - subseqIdx
}

func longestPalindrome(s string) int {
	ans := 0
	count := [256]int{}
	for _, c := range s {
		count[c]++
	}
	haveOdd := false

	for _, a := range count {
		if a == 0 {
			continue
		}

		if a%2 == 0 {
			ans += a
		} else {
			ans += a - 1
			haveOdd = true
		}
	}
	if haveOdd {
		return ans + 1
	}
	return ans
}
