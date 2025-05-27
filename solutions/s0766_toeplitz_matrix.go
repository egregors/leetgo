/*
	https://leetcode.com/problems/toeplitz-matrix/

	Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return
		false.

	A matrix is Toeplitz if every diagonal from top-left to bottom-right has the
		same elements.
*/

package solutions

func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if r > 0 && c > 0 && matrix[r-1][c-1] != matrix[r][c] {
				return false
			}
		}
	}
	return true
}
