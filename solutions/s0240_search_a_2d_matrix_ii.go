/*
	https://leetcode.com/problems/search-a-2d-matrix-ii/

	Write an efficient algorithm that searches for a target value in an m x n integer matrix.
	The matrix has the following properties:

		Integers in each row are sorted in ascending from left to right.
		Integers in each column are sorted in ascending from top to bottom.
*/

package solutions

func searchMatrix2(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	r, c := n-1, 0

	for r >= 0 && r < n && c >= 0 && c < m {
		if target == matrix[r][c] {
			return true
		} else if target < matrix[r][c] {
			r--
		} else if target > matrix[r][c] {
			c++
		}
	}

	return false
}
