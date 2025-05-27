/*
	https://leetcode.com/problems/transpose-matrix/

	Given a 2D integer array matrix, return the transpose of matrix.

	The transpose of a matrix is the matrix flipped over its main diagonal,
		switching the matrix's row and column indices.
*/

package solutions

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])

	res := make([][]int, m)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}
