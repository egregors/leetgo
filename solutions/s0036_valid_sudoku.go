/*
	https://leetcode.com/problems/valid-sudoku/

		Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
			validated according to the following rules:

		Each row must contain the digits 1-9 without repetition.
		Each column must contain the digits 1-9 without repetition.
		Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
			without repetition.

	Note:

		A Sudoku board (partially filled) could be valid but is not necessarily
			solvable.
		Only the filled cells need to be validated according to the mentioned rules.
*/

package solutions

const dot = '.'

func isValidSudoku(board [][]byte) bool {
	subBoxes := make(map[int][]byte, 9)
	rows := make(map[int][]byte, 9)
	cols := make(map[int][]byte, 9)

	for i := 0; i < len(board); i++ {
		rows[i] = board[i]
		for j := 0; j < len(board[i]); j++ {

			cols[j] = append(cols[j], board[i][j])
			if i < 3 {
				if j < 3 {
					subBoxes[1] = append(subBoxes[1], board[i][j])
				} else if j < 6 {
					subBoxes[2] = append(subBoxes[2], board[i][j])
				} else {
					subBoxes[3] = append(subBoxes[3], board[i][j])
				}
			} else if i < 6 {
				if j < 3 {
					subBoxes[4] = append(subBoxes[4], board[i][j])
				} else if j < 6 {
					subBoxes[5] = append(subBoxes[5], board[i][j])
				} else {
					subBoxes[6] = append(subBoxes[6], board[i][j])
				}
			} else {
				if j < 3 {
					subBoxes[7] = append(subBoxes[7], board[i][j])
				} else if j < 6 {
					subBoxes[8] = append(subBoxes[8], board[i][j])
				} else {
					subBoxes[9] = append(subBoxes[9], board[i][j])
				}
			}
		}
	}

	if !isMapValid(rows) || !isMapValid(cols) || !isMapValid(subBoxes) {
		return false
	}
	return true
}

func isMapValid(xs map[int][]byte) bool {
	for _, r := range xs {
		m := make(map[byte]bool)
		for _, e := range r {
			if e != dot {
				if m[e] {
					return false
				}
				m[e] = true
			}
		}
	}
	return true
}
