package main

import "fmt"

func canJump(nums []int) bool {
	farest := 0
	for i := 0; i <= farest; i++ {
		cur := i + nums[i]
		if cur > farest {
			farest = cur
		}

		if farest >= len(nums)-1 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(canJump([]int{0}))
}
