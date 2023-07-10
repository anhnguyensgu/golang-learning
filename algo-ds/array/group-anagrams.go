package main

import (
	"fmt"
	"sort"
)

func curCheck(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapChar := map[byte]int{}

	for i, _ := range s {
		mapChar[s[i]]++
		mapChar[t[i]]--
	}

	for _, c := range mapChar {
		if c != 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	mapAnagrams := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		curStr := ByRune(strs[i])
		sort.Sort(curStr)
		sortedStr := string(curStr)
		if _, ok := mapAnagrams[sortedStr]; ok {
			mapAnagrams[sortedStr] = append(mapAnagrams[sortedStr], strs[i])
		} else {
			mapAnagrams[sortedStr] = []string{strs[i]}
		}
	}

	ans := [][]string{}
	for _, a := range mapAnagrams {
		ans = append(ans, a)
	}
	return ans
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }
func main() {
	str := "cba"
	runes := ByRune(str)
	sort.Sort(runes)
	fmt.Println(string(runes))
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
