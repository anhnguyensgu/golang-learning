package main

func dfs(grid [][]int, row int, col int, seen [][]bool) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == 0 {
		return 0
	}

	seen[row][col] = true

	left := dfs(grid, row, col-1, seen)
	right := dfs(grid, row, col+1, seen)
	top := dfs(grid, row-1, col, seen)
	bot := dfs(grid, row+1, col, seen)

	return 1 + left + right + top + bot
}

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	seen := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		seen[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cur := dfs(grid, i, j, seen)
			if cur > ans {
				ans = cur
			}
		}
	}

	return ans
}

func main() {
	
}
