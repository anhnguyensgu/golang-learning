package productarrayexceptself

func productExceptSelf(nums []int) []int {
	zeros := 0
	total := 1
	for _, n := range nums {
		if n == 0 {
			zeros++
		} else {
			total *= n

		}
	}

	for i := range nums {
		if zeros == 1 {
			if nums[i] == 0 {
				nums[i] = total
			} else {
				nums[i] = 0
			}

		} else if zeros > 1 {
			nums[i] = 0
		} else {
			nums[i] /= total
		}
	}

	return nums
}
