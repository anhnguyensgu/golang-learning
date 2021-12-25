package main

import (
	"fmt"
	"reflect"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func checkWithHashMap(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var hashMap [128]int

	for _, c := range s1 {
		hashMap[c]++
	}

	for _, c := range s2 {
		hashMap[c]--
		if hashMap[c] != 0 {
			return false
		}
	}

	return true
}

func checkWithSort(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var runes1 ByRune = []rune(s1)
	var runes2 ByRune = []rune(s2)

	sort.Sort(runes1)
	sort.Sort(runes2)

	return reflect.DeepEqual(runes1, runes2)
}

func main() {
	s1 := "dog"
	s2 := "god"
	fmt.Println(checkWithHashMap(s1, s2))
}
