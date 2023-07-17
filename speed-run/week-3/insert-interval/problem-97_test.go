package insertinterval

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	intervals := [][]int{
		// {1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
		{1, 5},
	}
	newInterval := []int{2, 3}
	ans := insert(intervals, newInterval)
	fmt.Println(ans)
}
