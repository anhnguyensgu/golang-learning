package main

import "fmt"

func checkString(s string) bool {
	if len(s) == 1 {
		return true
	}
	i := -1
	j := -1
	for k:=0;k<len(s);k++ {
		if s[k] == 'a' {
			i = k
		} else if s[k] == 'b' {
			j = k
		}

		 if !(i == -1 || j == -1) && i > j {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkString("aaabbb"))
	fmt.Println(checkString("abab"))
	fmt.Println(checkString("babbbb"))
}
