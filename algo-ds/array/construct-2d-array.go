package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, 0)

	for i := 0; i < m*n; i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}

func main() {

	original := []int{1, 2, 3, 4}
	m := 2
	n := 2
	fmt.Println(construct2DArray(original, m, n))
}
