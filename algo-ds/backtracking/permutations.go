package main

func backtrack(nums []int, i int, ans [][]int) {
	if i == len(nums) {
		temp := make([]int, 0)
		for _, val := range nums {
			temp = append(temp, val)
		}
		ans = append(ans, temp)
		return
	}

	for index := i; index < len(nums); index++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(nums, i+1, ans)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func permute(nums []int) [][]int {
	var ans [][]int
	backtrack(nums, 0, ans)
	return ans
}

func main() {
	permute([]int{1, 2, 3})
}
