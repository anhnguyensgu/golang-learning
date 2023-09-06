package rottingoranges

const (
	Fresh  int = 1
	Rotten     = 2
)

var dir = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	freshCount := 0
	queue := [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == Fresh {
				freshCount++
			} else if grid[i][j] == Rotten {
				queue = append(queue, []int{i, j})
			}
		}

	}

	ans := 0
	for len(queue) != 0 {
		for _, orange := range queue {
			for _, d := range dir {
				x, y := orange[0]+d[0], orange[1]+d[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid) && grid[x][y] == Fresh {
					grid[x][y] = Rotten
					freshCount--
					queue = append(queue, []int{x, y})
				}

			}
		}
		ans++
		if freshCount == 0 {
			return ans
		}
	}

	return -1
}
