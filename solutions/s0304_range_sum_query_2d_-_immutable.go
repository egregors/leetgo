/*
	https://leetcode.com/problems/range-sum-query-2d-immutable/

	Given a 2D matrix matrix, handle multiple queries of the following type:

		Calculate the sum of the elements of matrix inside the rectangle defined by its upper left
		corner (row1, col1) and lower right corner (row2, col2).

	Implement the NumMatrix class:

		NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
		int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements
		of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower
		right corner (row2, col2).
*/

//nolint:revive // it's ok
package solutions

type NumMatrix struct {
	pref [][]int
}

// NewNumMatrix should call Constructor to pass LeetCode tests
func NewNumMatrix(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := len(matrix[0])

	pref := make([][]int, n)
	for r := 0; r < n; r++ {
		pref[r] = make([]int, m+1)
	}

	for r := 0; r < n; r++ {
		for c := 1; c <= m; c++ {
			pref[r][c] = pref[r][c-1] + matrix[r][c-1]
		}
	}

	return NumMatrix{pref: pref}
}

func (m *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		sum += m.pref[r][col2+1] - m.pref[r][col1]
	}
	return sum
}
