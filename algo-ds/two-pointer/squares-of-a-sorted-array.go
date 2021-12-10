package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	k := len(nums) - 1
	ans := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for i < j {
		if Abs(nums[i]) >= Abs(nums[j]) {
			ans[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			ans[k] = nums[j] * nums[j]
			k--
			j--
		}
	}
	return ans
}

func main() {
	nums := []int{2, 3, 6, 7}
	fmt.Print(sortedSquares(nums))
}
