package main

import (
	// "math"
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	count := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] > intervals[j][1]
	})

	currentRange := []int{-1, -1}

	for _, cur := range intervals {
		if cur[0] > currentRange[1] {
			count += 2
			currentRange[0] = cur[0] - 1
			currentRange[1] = cur[1]
		} else if cur[0] == currentRange[1] || cur[0] > currentRange[0] {
			count++
			currentRange[0] = currentRange[1]
			currentRange[1] = cur[1]
		}
	}
	return count
}

func main() {
	ans := intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}})
	fmt.Println(ans)
}
