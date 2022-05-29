package main

func rob(nums []int) int {
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

func main() {

}
