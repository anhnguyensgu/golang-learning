package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	prefix := []int{}

	last := 0
	for i := 0; i < len(nums); i++ {
		last += nums[i]
		prefix = append(prefix, last)
	}
	ans := prefix[3] - prefix[0]
	fmt.Println(ans)
}
