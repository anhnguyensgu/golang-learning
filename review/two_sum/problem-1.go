package two_sum

import "fmt"

// HashMap problem

func twoSum(nums []int, target int) []int {
	sumMap := map[int]int{}

	for idx, v := range nums {
		fmt.Println(target - v)
		if idx2, ok := sumMap[target-v]; ok {
			return []int{idx2, idx}
		} else {
			sumMap[v] = idx
		}
	}

	return []int{}
}
