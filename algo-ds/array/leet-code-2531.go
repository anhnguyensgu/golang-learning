package main

import "fmt"

func repeatedCharacter(s string) byte {
	charMap := [26]int{}
	for _, charValue := range s {
		if charMap[charValue-'a'] != 0 {
			return byte(charValue)
		}
		charMap[charValue-1]++
	}

	return (byte(1))

}

func main() {
	ans := repeatedCharacter("abc")
	fmt.Println(ans)
}
