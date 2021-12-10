package main

import "fmt"

const (
	Fresh  = 1
	Rotten = 2
)

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	queue := [][2]int{}
	freshCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case Fresh:
				freshCount++
			case Rotten:
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	ans := 0

	for len(queue) != 0 {
		for _, rotten := range queue {
			for _, direction := range directions {
				x, y := rotten[0]+direction[0], rotten[1]+direction[1]

				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == Fresh {
					grid[x][y] = Rotten
					freshCount--
					queue = append(queue, [2]int{x, y})
				}
			}

			queue = queue[1:]
		}

		ans++
		if freshCount == 0 {
			return ans
		}

	}

	return -1
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}
