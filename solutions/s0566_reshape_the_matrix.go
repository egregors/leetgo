/*
	https://leetcode.com/problems/reshape-the-matrix/

	In MATLAB, there is a handy function called reshape which can reshape an m x n matrix
	into a new one with a different size r x c keeping its original data.

	You are given an m x n matrix mat and two integers r and c representing the number
	of rows and the number of columns of the wanted reshaped matrix.

	The reshaped matrix should be filled with all the elements of the original
	matrix in the same row-traversing order as they were.

	If the reshape operation with given parameters is possible and legal, output the
	new reshaped matrix; Otherwise, output the original matrix.
*/

package solutions

func matrixReshape(mat [][]int, r, c int) [][]int {
	m := len(mat)
	n := len(mat[0])

	sum := m * n
	if sum != r*c {
		return mat
	}

	var res [][]int
	i := 0
	for a := 0; a < r && i < sum; a++ {
		row := make([]int, c)
		for b := 0; b < c && i < sum; b++ {
			row[b] = mat[i/n][i%n]
			i++
		}
		res = append(res, row)
	}

	return res
}
