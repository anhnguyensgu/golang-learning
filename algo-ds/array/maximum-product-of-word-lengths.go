package main

func maxProduct(words []string) int {
	ans := 0
	char := map[string]int{}
	for _, word := range words {
		bit := 0
		for _, c := range word {
			bit = bit | 1<<(c-'a')
		}
		char[word] = bit
	}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if char[words[i]]&char[words[j]] == 0 {
				cur := len(words[i]) * len(words[j])
				if cur > ans {
					ans = cur
				}
			}
		}
	}
	return ans
}

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	ans := maxProduct(words)
	println(ans)
}
