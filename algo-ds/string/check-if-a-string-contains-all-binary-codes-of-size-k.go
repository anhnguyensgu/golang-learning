package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	set := map[string]bool{}

	count := 1 << k
	fmt.Println(count)

	for i := 0; i < len(s)-k+1 && count != 0; i++ {
		cur := s[i : i+k]

		if set[cur] {
			continue
		}

		fmt.Printf("binary i %d %s\n", i, cur)

		set[cur] = true
		count--
	}

	return count == 0

}

func main() {
	s :=
		"00110"
	k := 2
	ans := hasAllCodes(s, k)
	fmt.Println(ans)
}
