package main

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	var left, right int64
	left, right = 0, 100000000000000
	for left < right {
		var mid int64 = (left + right) / 2
		var count int64 = 0
		for _, t := range time {
			a := mid / int64(t)
			count += a
		}

		if count < int64(totalTrips) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	// fmt.Println("aaaa")
}
