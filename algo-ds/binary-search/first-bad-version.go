package main

import "fmt"

/*
1	2	3	4	5
			x	x
*/

// because bad versions always exists => l < r is the condition to terminnate
// if mid is bad => maybe left handside has bad versions => assign right to mid
// else move left to mid + 1

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return n >= 1
}

func main() {
	fmt.Println(firstBadVersion(1))
}
