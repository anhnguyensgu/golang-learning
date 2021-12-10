package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	chars := []rune(s)
	i := 0
	j := len(chars) - 1

	for i < j {
		if !unicode.IsLetter(chars[i]) && !unicode.IsDigit(chars[i]) {
			i++
		} else if !unicode.IsLetter(chars[j]) && !unicode.IsDigit(chars[j]) {
			j--
		} else if unicode.ToLower(chars[i]) != unicode.ToLower(chars[j]) {
			return false
		} else {
			j--
			i++
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
