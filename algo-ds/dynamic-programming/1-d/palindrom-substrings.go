package main

import "fmt"

func isPalindrom(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func countSubstrings(s string) int {
	ans := 0
	for end := 0; end < len(s); end++ {
		for start := end; start >= 0; start-- {
			l, r := start, end
			fmt.Println(s[l : r+1])
			if isPalindrom(s, l, r) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	s := "aaa"
	ans := countSubstrings(s)
	fmt.Println(ans)
}
