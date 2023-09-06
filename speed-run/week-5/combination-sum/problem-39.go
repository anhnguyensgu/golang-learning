package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	dfs(candidates, target, &result, []int{}, 0)
	return result
}

func dfs(candidates []int, target int, result *[][]int, cur []int, pos int) {
	if target == 0 {
		c := []int{}
		c = append(c, cur...)
		*result = append(*result, c)
		return
	}

	for i := pos; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			continue
		}
		cur = append(cur, candidates[i])
		dfs(candidates, target-candidates[i], result, cur, i)
		cur = cur[:len(cur)-1]
	}
}
