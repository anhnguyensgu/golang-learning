package june

import (
	"fmt"
	"sort"
)

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

func commonChars(words []string) []string {
	set := [][26]int{}

	for i := range words {
		cur := [26]int{}
		for j := range words[i] {
			cur[words[i][j]-'a']++
		}
		set = append(set, cur)
	}

	ans := []string{}
	for c := 0; c < 26; c++ {
		times := 1000
		for i := range set {
			if times > set[i][c] {
				times = set[i][c]
			}
		}

		if times == 1000 {
			times = 0
		}

		for t := 0; t < times; t++ {
			ans = append(ans, fmt.Sprintf("%c", c+'a'))
		}

	}

	return ans
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := map[int]int{}
	for _, c := range hand {
		count[c]++
	}
	keys := []int{}
	for k := range count {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		for count[keys[i]] != 0 {
			last := keys[i]
			count[last]--
			k := i + 1

			for ; k < i+groupSize && k < len(keys); k++ {
				if keys[k]-last != 1 {
					return false
				}
				cur, ok := count[keys[k]]
				if !ok {
					return false
				}
				if cur < 1 {
					return false
				}
				count[keys[k]]--
				last = keys[k]
			}
			if k-i != groupSize {
				return false
			}
		}
	}

	return true

}
