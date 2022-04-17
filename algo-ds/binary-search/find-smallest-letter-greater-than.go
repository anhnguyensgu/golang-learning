package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters)
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return letters[l%len(letters)]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	var target byte
	target = 'z'
	fmt.Printf("%c", nextGreatestLetter(letters, target))
}
