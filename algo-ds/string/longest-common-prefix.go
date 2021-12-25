package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	i := 0
	ans := []rune{}

	isContinue := true
	for isContinue && i < len(strs[0]) {
		cur := strs[0][i]
		for _, str := range strs {
			if i >= len(str) || str[i] != cur {
				isContinue = false
				break
			}
		}
		i++
		if isContinue {
			ans = append(ans, rune(cur))
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
