/*
	https://leetcode.com/problems/set-matrix-zeroes/

	Given an m x n integer matrix matrix, if an element is 0, set its
	entire row and column to 0's.

	You must do it in place.
*/

package solutions

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	rs := make([]bool, n)
	cs := make([]bool, m)

	for ri, r := range matrix {
		for ci, c := range r {
			if c == 0 {
				rs[ri] = true
				cs[ci] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rs[i] || cs[j] {
				matrix[i][j] = 0
			}
		}
	}
}
