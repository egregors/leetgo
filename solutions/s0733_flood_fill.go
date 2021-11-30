/*
	https://leetcode.com/problems/flood-fill/

	An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

	You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from
	the pixel image[sr][sc].

	To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel
	of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same
	color), and so on. Replace the color of all of the aforementioned pixels with newColor.

	Return the modified image after performing the flood fill.
*/

package solutions

func fillDfs(image [][]int, color, newColor, r, c int) {
	if image[r][c] == color {
		image[r][c] = newColor
		if r >= 1 {
			fillDfs(image, color, newColor, r-1, c)
		}
		if r+1 < len(image) {
			fillDfs(image, color, newColor, r+1, c)
		}
		if c >= 1 {
			fillDfs(image, color, newColor, r, c-1)
		}
		if c+1 < len(image[0]) {
			fillDfs(image, color, newColor, r, c+1)
		}
	}
}

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	fillDfs(image, color, newColor, sr, sc)
	return image
}
