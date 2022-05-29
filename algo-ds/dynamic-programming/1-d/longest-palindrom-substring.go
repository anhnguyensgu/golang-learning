package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	l, r := 0, 0

	for start := len(s) - 1; start >= 0; start-- {
		for end := start + 1; end < len(s); end++ {
			if s[start] == s[end] {
				if dp[start+1][end-1] || start == end-1 {
					dp[start][end] = true

					if r-l <= end-start {
						l, r = start, end
					}
				}
			}
		}
	}

	return s[l : r+1]
}

func main() {
	s := "abcdedc"
	ans := longestPalindrome(s)
	fmt.Println(ans)

}
