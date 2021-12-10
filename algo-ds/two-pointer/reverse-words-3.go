package main

import "fmt"

func reverse(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	chars := []byte(s)
	j := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == ' ' {
			reverse(chars[j:i])
			j = i + 1
		}
	}
	if j < len(chars) {
		reverse(chars[j:len(chars)])
	}
	return string(chars)
}

func main() {
	fmt.Println(reverseWords("s'teL ekat edoCteeL tsetnoc"))
}
