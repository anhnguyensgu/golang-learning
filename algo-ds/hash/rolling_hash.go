package main

import "fmt"

func RollingHash(s string) int {
	rh := make([]int, 26)

	M := 1000000007

	rh[0] = 26
	for i := 1; i < 26; i++ {
		rh[i] = (rh[i-1] * 26) % M
	}

	val := 0
	for _, r := range s {
		var ind = int(r - 'a')
		val = (val + int(r)*rh[ind]) % M
	}
	return val
}
func main() {
	str := "abc"
	str2 := "cba"

	fmt.Println(RollingHash(str))
	fmt.Println(RollingHash(str2))
}
