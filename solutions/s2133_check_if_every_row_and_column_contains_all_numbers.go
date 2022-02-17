/*
	https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/

	An n x n matrix is valid if every row and every column contains all the integers from 1
	to n (inclusive).

	Given an n x n integer matrix matrix, return true if the matrix is valid. Otherwise, return
	false.
*/

package solutions

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	rowM := make([]map[int]int, n)
	colM := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		rowM[i] = make(map[int]int)
		colM[i] = make(map[int]int)
	}

	for ri := 0; ri < n; ri++ {
		for ci := 0; ci < n; ci++ {
			rowM[ri][matrix[ri][ci]]++
			colM[ci][matrix[ri][ci]]++
		}
		if len(rowM[ri]) != n {
			return false
		}
	}
	for _, m := range colM {
		if len(m) != n {
			return false
		}
	}

	return true
}
