package main

import (
	"fmt"
	"sort"
)

type Matrix [][]int

func (r Matrix) Len() int           { return len(r) }
func (r Matrix) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Matrix) Less(i, j int) bool { return r[i][0] < r[j][0] }

func removeCoveredIntervals(intervals [][]int) int {
	var a Matrix = intervals
	sort.Sort(a)

	ans := 1
	x1 := a[0][0]
	x2 := a[0][1]
	for _, interval := range a {
		if x2 < interval[1] && x1 < interval[0] {
			ans++
		}

		if x2 < interval[1] {
			x2 = interval[1]
			x1 = interval[0]
		}
	}

	return ans
}

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{2, 9}, {3, 6}, {2, 8}}))
}
