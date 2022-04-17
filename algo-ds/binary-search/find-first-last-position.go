package main

import "fmt"

func findFirst(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	var ans int
	if ans = l; l >= len(nums) || nums[l] != target {
		ans = -1
	}
	return ans
}

func findLast(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	ans := r
	if r < 0 || nums[r] != target {
		ans = -1
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	ans[0] = findFirst(nums, target)
	ans[1] = findLast(nums, target)
	return ans
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
}
