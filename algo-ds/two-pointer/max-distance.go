package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	ans := 0
	var i = 0
	var j = 0

	for i < len(nums1) && j < len(nums2) {
		if nums2[j] >= nums1[i] {
			newPos := j - i
			if ans < newPos {
				ans = newPos
			}
			j++
		} else {
			i++
		}
	}

	return ans
}

func main() {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}
	fmt.Println(maxDistance(nums1, nums2))
}
