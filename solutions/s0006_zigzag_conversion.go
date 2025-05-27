/*
	https://leetcode.com/problems/zigzag-conversion/

	The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of
		rows like this:
	(you may want to display this pattern in a fixed font for better legibility)
*/

package solutions

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	field := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		field[i] = make([]rune, len(s))
	}

	col, row, i := 0, 0, 0

	for i < len(s) {
		for row < numRows && i < len(s) {
			field[row][col] = rune(s[i])
			row++
			i++
		}
		row -= 2
		col++
		for row >= 0 && i < len(s) {
			field[row][col] = rune(s[i])
			row--
			col++
			i++
		}
		row += 2
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if field[i][j] != rune(0) {
				res += string(field[i][j])
			}
		}
	}

	return res
}
