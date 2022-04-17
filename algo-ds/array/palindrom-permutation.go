package main

import "fmt"

func isPalindromePermutation(s string) bool {
	ans := true
	charMap := make([]int, 26)

	for _, value := range s {
		if value-'a' < 0 {
			continue
		}
		charMap[value-'a'] = charMap[value-'a'] + 1
	}
	foundOdd := false
	fmt.Println(charMap)
	for _, value := range charMap {
		if value%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}

	}
	return ans
}

/*
using bitVector to reduce space
0 0 0 0 0
^
0 0 0 0 1
= 0 0 0 0 1
at the end, need to check if the bit string has at most 1's bit

0 1 0 1 0 0  - 1 = 0 1 0 0 1 1

0 1 0 1 0 0
&
0 1 0 0 1 1

*/
func bitVector(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		curChar := s[i] - 'a'
		if curChar >= 0 {
			count = count ^ 1<<curChar
		}
	}
	fmt.Println(count)
	ans := count & (count - 1)
	return ans == 0
}

func reduceLoop(s string) bool {
	charMap := make([]int, 26)
	count := 0
	for _, value := range s {
		curChar := value - 'a'

		if curChar < 0 {
			continue
		}
		charMap[curChar] = charMap[curChar] + 1
		if charMap[curChar]%2 == 1 {
			count--
		} else {
			count++
		}
	}
	fmt.Println(count)
	return count >= -1
}

func main() {
	fmt.Println(bitVector("tct   ca"))
	fmt.Println(bitVector("tct   a"))
}
