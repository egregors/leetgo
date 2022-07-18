/*
	https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

	Given a matrix and a target, return the number of non-empty submatrices that sum to target.

	A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.

	Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that
	is different: for example, if x1 != x1'.
*/

package solutions

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])
	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, m+1)
	}

	// 2d pref
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pref[i][j] = matrix[i-1][j-1] + pref[i-1][j] + pref[i][j-1] - pref[i-1][j-1]
		}
	}

	var count, sum int

	// 1d pref
	for c1 := 1; c1 <= m; c1++ {
		for c2 := c1; c2 <= m; c2++ {
			h := make(map[int]int)
			h[0] = 1
			for row := 1; row <= n; row++ {
				sum = pref[row][c2] - pref[row][c1-1]

				if v, ok := h[sum-target]; ok {
					count += v
				}

				if v, ok := h[sum]; ok {
					h[sum] = v + 1
				} else {
					h[sum] = 1
				}
			}
		}
	}

	return count
}
