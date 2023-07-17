package ransomenote

func canConstruct(ransomNote string, magazine string) bool {
	mapChar := [26]int{}

	for i := 0; i < len(mapChar); i++ {
		mapChar[i] = -1
	}
	for _, a := range ransomNote {
		if mapChar[a-'a'] == -1 {
			mapChar[a-'a'] = 1
		} else {
			mapChar[a-'a']++
		}
	}

	for _, a := range magazine {
		mapChar[a-'a']--
	}

	for _, v := range mapChar {
		if v > 0 || v < -1 {
			return false
		}
	}

	return true
}
