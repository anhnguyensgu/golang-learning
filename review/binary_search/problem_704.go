package binarysearch

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := r - (r-l)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
