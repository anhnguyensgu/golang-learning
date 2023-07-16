package floodfill

func dfs(image [][]int, r, c, targetColor, newColor int) {
	if image[r][c] == targetColor {
		image[r][c] = newColor
		if r > 0 {
			dfs(image, r-1, c, targetColor, newColor)
		}
		if r < len(image)-1 {
			dfs(image, r+1, c, targetColor, newColor)
		}

		if c > 0 {
			dfs(image, r, c-1, targetColor, newColor)
		}

		if c < len(image[0])-1 {
			dfs(image, r, c+1, targetColor, newColor)
		}
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]
	if curColor == newColor {
		return image
	}
	dfs(image, sr, sc, curColor, newColor)
	return image
}
