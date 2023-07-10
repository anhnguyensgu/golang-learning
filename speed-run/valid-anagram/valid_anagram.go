package validanagram

func isAnagram(s string, t string) bool {
	sl := len(s)
	tl := len(t)

	if sl != tl {
		return false
	}

	charMap := [26]int{}
	for i := 0; i < sl; i++ {
		charMap[s[i]-'a']++
		charMap[t[i]-'a']--
	}

	for _, c := range charMap {
		if c != 0 {
			return false
		}
	}

	return true
}
