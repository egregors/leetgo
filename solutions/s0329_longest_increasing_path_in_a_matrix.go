/*
	https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

	Given an m x n integers matrix, return the length of the longest increasing path in matrix.

	From each cell, you can either move in four directions: left, right, up, or down.
	You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
*/

package solutions

func longestIncreasingPath(matrix [][]int) int {
	cache := make([][]int, len(matrix))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}

	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans = Maximum(ans, dfs329(matrix, i, j, cache))
		}
	}
	return ans
}

func dfs329(m [][]int, i, j int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	dirs := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for _, d := range dirs {
		x, y := i+d[0], j+d[1]
		if x >= 0 && x < len(m) && y >= 0 && y < len(m[0]) && m[x][y] > m[i][j] {
			cache[i][j] = Maximum(cache[i][j], dfs329(m, x, y, cache))
		}
	}
	cache[i][j]++
	return cache[i][j]
}
