package main

import "fmt"

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	chars := []rune(palindrome)
	limit := len(chars) / 2
	for i := 0; i < limit; i++ {
		if chars[i] == 'a' {
			continue
		}
		chars[i] = 'a'
		return string(chars)
	}
	chars[len(chars)-1] = 'b'
	return string(chars)
}

func main() {
	palindrome := "aba"

	ans := breakPalindrome(palindrome)

	fmt.Println(ans)
}
