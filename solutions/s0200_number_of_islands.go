/*
	https://leetcode.com/problems/number-of-islands/

	Given an m x n 2D binary grid grid which represents a map of '1's (land) and
		'0's (water), return the number of islands.

	An island is surrounded by water and is formed by connecting adjacent lands
		horizontally or
	vertically. You may assume all four edges of the grid are all surrounded by
		water.
*/

package solutions

func dfs200(r, c int, grid [][]byte, seen map[[2]int]bool) bool {
	if seen[[2]int{r, c}] {
		return false
	}

	seen[[2]int{r, c}] = true
	if grid[r][c] == '0' {
		return false
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, d := range dirs {
		nR, nC := r+d[0], c+d[1]
		if nR >= 0 && nR < len(grid) && nC >= 0 && nC < len(grid[0]) {
			dfs200(nR, nC, grid, seen)
		}
	}

	return true
}

func numIslands(grid [][]byte) int {
	islands := 0

	seen := make(map[[2]int]bool)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if dfs200(r, c, grid, seen) {
				islands++
			}
		}
	}

	return islands
}
