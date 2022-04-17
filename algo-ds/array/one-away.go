package main

import (
	"fmt"
)

/*
remove = insert by swap s1 s2
insert abc ac = remove ac abc

remove/insert
if a[i] != b[j]
case 1 first found => increase i
case 2 second found => return false
else
increase i, j

replace
if a[i] != b[j]
case first found => increase i, j
case second found => return false
else
increase i, j
*/

func check(first string, second string) bool {
	ans := true

	var s1 string
	var s2 string
	diff := len(first) - len(second)
	if diff == 1 {
		s1 = first
		s2 = second
	} else if diff == -1 {
		s1 = second
		s2 = first
	} else {
		s1 = first
		s2 = second
	}
	i := 0
	j := 0
	foundDiff := false
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			if foundDiff {
				return false
			}
			foundDiff = true

			if diff == 0 {
				j++
			}
		} else {
			j++
		}
		i++
	}

	return ans
}

func main() {
	fmt.Println(check("pale", "ple"))
	fmt.Println(check("pales", "pale"))
	fmt.Println(check("pale", "bale"))
	fmt.Println(check("pale", "bake"))
}
