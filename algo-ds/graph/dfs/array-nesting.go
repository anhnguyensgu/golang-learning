package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	visted := make([]bool, len(nums))
	ans := 0

	for i := 0; i < len(nums); i++ {
		cur := i
		count := 0
		for !visted[cur] {
			visted[cur] = true
			cur = nums[cur]
			count++
		}
		if count > ans {
			ans = count
		}
	}

	fmt.Println(ans)
}
