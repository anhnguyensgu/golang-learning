package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	var mapCount [128]int

	for _, c := range s1 {
		mapCount[c]++
	}

	j, i := 0, 0

	for k := 0; k < len(s2); k++ {
		c := s2[k]
		i++
		if mapCount[c] > 0 {
			mapCount[c]--
			if i-j == len(s1) {
				return true
			}
		} else {
			for s2[j] != c {
				mapCount[s2[j]]++
				j++
			}
			j++
			// mapCount[s2[j]]++
			// j++
		}
	}

	return false
}

func main() {
	// fmt.Println(checkInclusion("ky",
	// 	"ainwkckifykxlribaypk"))
	fmt.Println(checkInclusion("ab",
		"eidboaoo"))
}
