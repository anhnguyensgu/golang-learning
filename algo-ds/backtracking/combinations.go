package main

import "fmt"

var ans [][]int

func helper(n int, k int, cur []int) {
	if len(cur) == k {
		temp := []int{}
		for _, num := range cur {
			temp = append(temp, num)
		}
		ans = append(ans, temp)
		return
	}

	// i will at most be n - (k - c.size()) + 1
	for i := cur[len(cur)-1] + 1; i <= n-(k-len(cur))+1; i++ {
		cur = append(cur, i)
		helper(n, k, cur)
		cur = cur[0 : len(cur)-1]
	}
}

func combine(n int, k int) [][]int {

	cur := []int{}
	ans = [][]int{}
	for i := 1; i <= n; i++ {
		cur = append(cur, i)
		helper(n, k, cur)
		cur = cur[0 : len(cur)-1]
	}
	return ans
}

func main() {
	fmt.Println(combine(5, 4))
}
