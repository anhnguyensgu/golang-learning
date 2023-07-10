package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ans := true
	charMap := [26]int{}
	for _, c := range magazine {
		charMap[c-'a']++
	}

	for _, c := range ransomNote {
		charMap[c-'a']--
	}

	for _, c := range charMap {
		if c < 0 {
			ans = false
			break
		}
	}
	return ans
}

func main() {
	ransomeNote := "aaaaa"
	magazine := "aab"
	ans := canConstruct(ransomeNote, magazine)
	fmt.Print(ans)
}
