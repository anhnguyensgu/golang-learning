package main

import "fmt"

func wiggleMaxLengthNthSpace(nums []int) int {
	up := make([]int, len(nums))
	down := make([]int, len(nums))
	up[0] = 1
	down[0] = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i] = up[i-1] + 1
			up[i] = up[i-1]
		} else {
			down[i] = down[i-1]
			up[i] = up[i-1]
		}
	}

	maxUp := up[len(up)-1]
	maxDown := down[len(down)-1]
	if maxUp > maxDown {
		return maxUp
	} else {
		return maxDown
	}
}

func wiggleMaxLength(nums []int) int {
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	} else {
		return down
	}
}
func main() {
	nums := []int{1, 3, 1, 3}
	ans := wiggleMaxLengthNthSpace(nums)
	fmt.Println(ans)
}
