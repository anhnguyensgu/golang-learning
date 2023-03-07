package main

import (
	"fmt"
	"math"
	"sort"
)

//a < b < c < d
// for a
// 4a = a + a + a + a = a + b + c + d + (a-b) + (a-c) + (a-d) =
// for b
// 3b = b + c + d + ans

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	sum := int64(0)
	for _, a := range beans {
		sum += int64(a)
	}

	ans := int64(math.MaxInt)
	last := int64(0)
	for i, a := range beans {
		s := int64((len(beans) - i) * a)
		cur := sum - s
		if cur < ans {
			ans = cur
		}
		last += int64(a)
	}

	return ans
}

func main() {
	fmt.Println(minimumRemoval([]int{4, 1, 6, 5}))
}
