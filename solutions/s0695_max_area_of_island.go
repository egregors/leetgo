/*
	https://leetcode.com/problems/max-area-of-island/

	You are given an m x n binary matrix grid. An island is a group of 1's
		(representing land)
	connected 4-directionally (horizontal or vertical.) You may assume all four
		edges of the grid are
	surrounded by water.

	The area of an island is the number of cells with a value 1 in the island.

	Return the maximum area of an island in grid. If there is no island, return 0.
*/

package solutions

import "fmt"

func countIsland(r, c int, grid [][]int, seen map[string]bool) int {
	size := 0
	if grid[r][c] == 1 && !seen[fmt.Sprintf("%dx%d", c, r)] {
		size++
		seen[fmt.Sprintf("%dx%d", c, r)] = true
		if r >= 1 {
			size += countIsland(r-1, c, grid, seen)
		}
		if r+1 < len(grid) {
			size += countIsland(r+1, c, grid, seen)
		}
		if c >= 1 {
			size += countIsland(r, c-1, grid, seen)
		}
		if c+1 < len(grid[0]) {
			size += countIsland(r, c+1, grid, seen)
		}
	}

	return size
}

func maxAreaOfIsland(grid [][]int) int {
	seen := make(map[string]bool)
	max := 0

	for ci := 0; ci < len(grid[0]); ci++ {
		for ri := 0; ri < len(grid); ri++ {
			s := countIsland(ri, ci, grid, seen)
			if s > max {
				max = s
			}
		}
	}
	return max
}
