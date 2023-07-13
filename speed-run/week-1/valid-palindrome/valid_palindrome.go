package validpalindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	chars := []rune(s)

	for l < r {
		if !unicode.IsLetter(chars[l]) && !unicode.IsDigit(chars[l]) {
			l++
		} else if !unicode.IsLetter(chars[r]) && !unicode.IsDigit(chars[r]) {
			r--
		} else if unicode.ToLower(chars[l]) != unicode.ToLower(chars[r]) {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
