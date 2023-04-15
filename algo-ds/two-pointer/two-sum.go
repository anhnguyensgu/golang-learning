package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}

	}
	return []int{}
}

func main() {
	nums := []int{0, 1, 0, 2, 3}
	fmt.Println(twoSum(nums, 3))
}
