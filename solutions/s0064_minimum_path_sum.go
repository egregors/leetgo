/*
	https://leetcode.com/problems/minimum-path-sum/

	Given a m x n grid filled with non-negative numbers, find a path from top left
		to bottom right,
	which minimizes the sum of all numbers along its path.

	Note: You can only move either down or right at any point in time.
*/

package solutions

func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += Minimum(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
