package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return l
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7

	fmt.Println(searchInsert(nums, target))
}
