package main

import (
	"fmt"
	"sort"
)

type Matrix [][]int

func (r Matrix) Len() int           { return len(r) }
func (r Matrix) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Matrix) Less(i, j int) bool { return r[i][0] < r[j][0] }

func merge(intervals [][]int) [][]int {
	var ans Matrix = [][]int{}
	var sortedInternal Matrix = intervals
	sort.Sort(sortedInternal)

	for i := 0; i < len(intervals); i++ {
		if len(ans) == 0 || ans[len(ans) - 1][1] < sortedInternal[i][0] {
			ans = append(ans, sortedInternal[i])
		} else {
			end := sortedInternal[i][1]
			if end < ans[len(ans) - 1][1] {
				end = ans[len(ans) - 1][1]
			}
			ans[len(ans)-1][1] = end
		}
	}

	return ans
}

func main() {
	intervals := [][]int{
		{0, 5},
		{8, 10},
		{1, 4},
		{0, 18},
		{15, 18},
	}

	fmt.Println(merge(intervals))
}
