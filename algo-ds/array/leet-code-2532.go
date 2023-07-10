package main

import "fmt"

func equalPairs(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count := 0
			for k := 0; k < len(grid); k++ {
				if grid[i][k] == grid[k][j] {
					count++
				}
			}
			if count == len(grid) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	gird := [][]int{}

	ans := equalPairs(gird)
	fmt.Println(ans)
}
