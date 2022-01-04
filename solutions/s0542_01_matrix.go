/*
	https://leetcode.com/problems/01-matrix/

	Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

	The distance between two adjacent cells is 1.
*/

package solutions

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	res := make([][]int, rows)
	for r := 0; r < rows; r++ {
		res[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			res[r][c] = 999999
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				res[r][c] = 0
			} else {
				if r > 0 {
					res[r][c] = min(res[r][c], res[r-1][c]+1)
				}
				if c > 0 {
					res[r][c] = min(res[r][c], res[r][c-1]+1)
				}
			}
		}
	}

	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			if r < rows-1 {
				res[r][c] = min(res[r][c], res[r+1][c]+1)
			}
			if c < cols-1 {
				res[r][c] = min(res[r][c], res[r][c+1]+1)
			}
		}
	}

	return res
}
