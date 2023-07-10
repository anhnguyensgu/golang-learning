package main

import (
	"fmt"
	"strconv"
)

var cache = map[int]string{}

func say(term string) string {
	count := 0
	cur := term[0]
	result := ""
	for i := 0; i < len(term); i++ {
		if term[i] == cur {
			count++
		} else {
			result = result + strconv.Itoa(count) + strconv.Itoa(int(cur-'0'))
			count = 1
			cur = term[i]
		}
	}
	result = result + strconv.Itoa(count) + strconv.Itoa(int(cur-'0'))

	return result
}

func countAndSay(n int) string {
	ans := "1"
	for i := 1; i <= n-1; i++ {
		if _, ok := cache[i]; ok {
			ans = cache[i]
			continue
		}
		ans = say(ans)
		cache[i] = ans
	}
	return ans
}

func main() {
	//n := 0
	fmt.Print(countAndSay(4))
}
