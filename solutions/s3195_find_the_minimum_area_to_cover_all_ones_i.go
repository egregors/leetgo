/*
	https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/

	You are given a 2D binary array grid. Find a rectangle with horizontal and
		vertical sides with the smallest area, such that all the 1's in grid lie
		inside this rectangle.

	Return the minimum possible area of the rectangle.
*/

package solutions

func minimumArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	minI, maxI := n, 0
	minJ, maxJ := m, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				minI = min(minI, i)
				maxI = max(maxI, i)
				minJ = min(minJ, j)
				maxJ = max(maxJ, j)
			}
		}
	}
	return (maxI - minI + 1) * (maxJ - minJ + 1)
}
