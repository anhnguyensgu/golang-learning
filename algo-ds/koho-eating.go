package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	fmt.Printf("max %d min %d\n", max, min)

	ans := max
	for min <= max {
		mid := min + (max-min)/2
		sum := 0
		for _, p := range piles {

			if p%mid != 0 {
				sum += (p / mid) + 1
			} else {
				sum += (p / mid)
			}
		}
		remainingHour := h - sum
		fmt.Printf("k :%d remainingHour: %d sum: %d\n", mid, remainingHour, sum)
		if remainingHour < 0 {
			min = mid + 1
		} else {
			if ans > mid {
				ans = mid
			}
			max = mid - 1
		}
	}
	return ans

}

func main() {
	piles := []int{312884470}
	h := 312884469
	fmt.Println(minEatingSpeed(piles, h))
}
