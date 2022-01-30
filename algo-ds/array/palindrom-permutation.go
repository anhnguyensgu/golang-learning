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
	fmt.Println(reduceLoop("tct   a"))
}
