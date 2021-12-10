package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	ans := 0
	var mapCount [128]int

	j := 0
	var i int
	for i = 0; i < len(s); i++ {
		for mapCount[s[i]] != 0 {
			// fmt.Printf("change %d %d \n", i, j)
			mapCount[s[j]]--
			j++
		}
		mapCount[s[i]]++
		if i-j+1 > ans {
			// fmt.Printf("%d %d \n", i, j)
			ans = i - j + 1
		}

	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("bcaaaabcd"))
}
