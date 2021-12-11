package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	dist := [][]int{}

	q := [][2]int{}

	for i, row := range mat {
		temp := []int{}
		for j, cell := range row {
			if cell == 0 {
				temp = append(temp, 0)
				q = append(q, [2]int{i, j})
			} else {
				temp = append(temp, 1e10)
			}
		}
		dist = append(dist, temp)
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) != 0 {
		cur := q[0]
		for _, dir := range dirs {
			x, y := cur[0]+dir[0], cur[1]+dir[1]

			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) {
				if dist[x][y] > dist[cur[0]][cur[1]]+1 {
					dist[x][y] = dist[cur[0]][cur[1]] + 1
					q = append(q, [2]int{x, y})
				}
			}
		}
		q = q[1:]

	}
	return dist
}

//review later, fucking lazy.
func updateMatrixDP(mat [][]int) [][]int {
	return nil
}

func main() {
	mat := [][]int{
		{1, 1, 0},
		// {0, 1, 0},
		// {1, 1, 1},
	}

	fmt.Println(updateMatrix(mat))
}
