package subsets

func subsets(nums []int) [][]int {
	ans := [][]int{}
	cur := []int{}
	ans = append(ans, []int{})
	for i := 1; i <= len(nums); i++ {
		dfs(&ans, nums, cur, i, 0)
	}

	return ans
}

func dfs(ans *[][]int, nums []int, cur []int, size, curInd int) {
	if len(cur) > size {
		return
	}

	if len(cur) == size {
		a := []int{}
		for _, i := range cur {
			a = append(a, i)
		}
		*ans = append(*ans, a)
	}

	for i := curInd; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfs(ans, nums, cur, size, i+1)
		temp := cur
		cur = temp[0 : len(temp)-1]
	}
}
