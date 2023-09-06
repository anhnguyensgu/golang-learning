package numberofislands

func numIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans += 1
				dfs(i, j, grid)
			}
		}
	}
	return ans
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j > len(grid[0]) {
		return
	}

	if grid[i][j] == '1' {
		// grid[i][j] = 0
		dfs(i-1, j, grid)
		dfs(i+1, j, grid)
		dfs(i, j-1, grid)
		dfs(i, j+1, grid)
	}
}
