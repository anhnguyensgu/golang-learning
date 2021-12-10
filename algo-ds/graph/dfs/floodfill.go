package main

import "fmt"

func dfs(image [][]int, sr int, sc int, newColor int, color int) [][]int {
	rowLength := len(image)
	columnLength := len(image[0])
	if sr < 0 || sr >= rowLength || sc < 0 || sc >= columnLength {
		return image
	}

	if image[sr][sc] == color {
		image[sr][sc] = newColor
		dfs(image, sr-1, sc, newColor, color) // move up
		dfs(image, sr+1, sc, newColor, color) // move down
		dfs(image, sr, sc-1, newColor, color) // move left
		dfs(image, sr, sc+1, newColor, color) // move right
	}
	return image
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if image[sr][sc] != newColor {
		image[sr][sc] = newColor
		dfs(image, sr-1, sc, newColor, color) // move up
		dfs(image, sr+1, sc, newColor, color) // move down
		dfs(image, sr, sc-1, newColor, color) // move left
		dfs(image, sr, sc+1, newColor, color) // move right
	}

	return image
}

func main() {
	image := [][]int{
		{0, 0, 0},
		{0, 1, 0},
	}

	floodFill(image, 1, 1, 2)

	fmt.Println(image)
}
