/*
	https://leetcode.com/problems/island-perimeter/description

	You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

	Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
	and there is exactly one island (i.e., one or more connected land cells).

	The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island.
	One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100.
	Determine the perimeter of the island.
*/

package solutions

import "fmt"

func islandPerimeter(grid [][]int) int {
	m := make(map[string]bool)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				m[fmt.Sprintf("%d:%d", r, c)] = true
			}
		}
	}

	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	res := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
						if m[fmt.Sprintf("%d:%d", nr, nc)] {
							continue
						}
						res++
						continue
					}
					res++
				}
			}
		}
	}

	return res
}
