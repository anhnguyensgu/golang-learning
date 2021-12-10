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

// 1 2 3 4 5 6 7
// [4] 2 3 [1] 5 6 7
// [4  5] 3 [1 2 6 7]
// [4 5 6] [1 2 3 7]

func rotateSwapKSize(nums []int, k int) []int {
	size := len(nums)
	k = k % size

	low := 0
	step := k

	for low < size && step > 0 {
		for i := 0; i < step; i++ {
			nums[low+i], nums[size-step+i] = nums[size-step+i], nums[low+i]
		}
		low = low + step
		step = step % (size - low)
	}
	return nums
}

func rotate(nums []int, k int) []int {
	return rotateSwapKSize(nums, k)
}

func main() {
	nums := []int{2, 3, 6, 7}
	k := 3
	fmt.Print(rotate(nums, k))
}
