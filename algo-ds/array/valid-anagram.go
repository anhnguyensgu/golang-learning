package main

import "fmt"

func isAnagram(s string, t string) bool {
	mapChar := map[rune]int{}

	for _, c := range s {
		mapChar[c]++
	}

	for _, c := range s {
		mapChar[c]--
	}

	for _, c := range mapChar {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "rat"
	t := "cat"
	ans := isAnagram(s, t)
	fmt.Println(ans)
}
