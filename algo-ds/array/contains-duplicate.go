package main

import "sort"

func containsDuplicate(nums []int) bool {
	hashSet := map[int]bool{}

	for _, num := range nums {
		if _, ok := hashSet[num]; ok {
			return true
		} else {
			hashSet[num] = true
		}
	}
	return false
}

func containsDuplicateUsingXor(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1]^nums[i] == 0 {
			return true
		}
	}
	return false
}
