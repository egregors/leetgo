/*
	https://leetcode.com/problems/unique-paths-ii/

	You are given an m x n integer array grid. There is a robot initially located
		at the top-left corner (i.e., grid[0][0]).
	The robot tries to move to the bottom-right corner (i.e., grid[m-1][n-1]). The
		robot can only move either down or right
	at any point in time.

	An obstacle and space are marked as 1 or 0 respectively in grid. A path that
		the robot takes cannot include any square
	that is an obstacle.

	Return the number of possible unique paths that the robot can take to reach the
		bottom-right corner.

	The testcases are generated so that the answer will be less than or equal to 2
		* 109.
*/

package solutions

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// first column
	for j := 1; j < n; j++ {
		if dp[0][j-1] != 0 && obstacleGrid[0][j] != 1 {
			dp[0][j] = 1
		}
	}
	// first row
	for i := 1; i < m; i++ {
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
