package main

import "fmt"

func minAddToMakeValid(s string) int {
	ans := 0
	c := 0

	for _, r := range s {
		if r == '(' {
			c++
		} else if c > 0 {
			c--
		} else {
			ans++
		}
	}

	if c > 0 {
		ans += c
	}

	return ans
}

func main() {
	fmt.Println(minAddToMakeValid(")))"))
	fmt.Println(minAddToMakeValid("()())"))
	fmt.Println(minAddToMakeValid("()"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()))(("))
}
