package main

import (
	"fmt"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func bruteforce(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}

			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func hashMapApproach(s string) bool {
	mapCount := make([]rune, 1000)
	for _, c := range s {
		if mapCount[c] != 0 {
			return false
		}
		mapCount[c]++
	}
	return true
}

func sortApproach(s string) bool {
	var chars ByRune = []rune(s)
	sort.Sort(chars)
	fmt.Println(chars)

	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			return false
		}
	}
	return true
}

func bitWise(s string) bool {
	checker := 0
	for _, c := range s {
		val := int(c) - 'a'
		if checker&(1<<val) > 0 {
			return false
		}
		checker = checker | 1<<val
	}
	return true
}

func main() {
	a := "abc"
	fmt.Println(bitWise(a))
}
