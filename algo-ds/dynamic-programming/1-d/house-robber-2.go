package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans1 := rob1(nums[1:])
	ans2 := rob1(nums[:len(nums)-1])
	return max(ans1, ans2)
}

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	var second int
	if nums[0] < nums[1] {
		second = nums[1]
	} else {
		second = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		temp := second
		if first+nums[i] > temp {
			temp = first + nums[i]
		}

		first = second
		second = temp
	}
	return second
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{2, 3, 2}
	ans1 := rob1(nums[1:])
	ans2 := rob1(nums[0 : len(nums)-1])
	fmt.Printf("%d", max(ans1, ans2))
}
