package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	validMap := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else if c == '}' || c == ']' || c == ')' {
			fmt.Println(stack)
			l := len(stack) - 1
			if l < 0 || validMap[c] != stack[l] {
				return false
			}
			stack = stack[0:l]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
