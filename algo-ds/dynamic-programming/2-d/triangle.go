package main

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	for row := n - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}

	return triangle[0][0]
}

func main() {

}
