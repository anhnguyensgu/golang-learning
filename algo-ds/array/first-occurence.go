package main

import "fmt"

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}

			if j == n-1 {
				return i
			}

		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("aaaaleeto", "leeto"))
}
