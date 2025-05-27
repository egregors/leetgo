/*
	https://leetcode.com/problems/out-of-boundary-paths/

	There is an m x n grid with a ball. The ball is initially at the position
	[startRow, startColumn]. You are allowed to move the ball to one of the four
	adjacent cells in the grid (possibly out of the grid crossing the grid
		boundary).
	You can apply at most maxMove moves to the ball.

	Given the five integers m, n, maxMove, startRow, startColumn, return the number
		of
	paths to move the ball out of the grid boundary. Since the answer can be very
		large,
	return it modulo 10^9 + 7.
*/

package solutions

import "fmt"

func findPaths(m, n, maxMove, startRow, startColumn int) int {
	seen := make(map[string]int)
	var bt func(r, c, steps int) int
	bt = func(r, c, steps int) int {
		if v, ok := seen[fmt.Sprintf("%d_%d_%d", r, c, steps)]; ok {
			return v
		}

		if r < 0 || r >= m || c < 0 || c >= n {
			return 1
		}
		if steps == 0 {
			return 0
		}

		paths := 0
		dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for _, d := range dirs {
			newR, newC := r+d[0], c+d[1]
			paths += bt(newR, newC, steps-1)
		}
		paths %= 1000000007
		seen[fmt.Sprintf("%d_%d_%d", r, c, steps)] = paths
		return paths
	}

	return bt(startRow, startColumn, maxMove)
}
