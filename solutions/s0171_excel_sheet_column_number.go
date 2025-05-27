/*
	https://leetcode.com/problems/excel-sheet-column-number/

	Given a string columnTitle that represents the column title as appear in an
		Excel sheet,
	return its corresponding column number.
*/

package solutions

func titleToNumber(columnTitle string) int {
	acc := 0
	for i := 0; i < len(columnTitle); i++ {
		acc *= 26
		acc += int(columnTitle[i] - 'A' + 1)
	}
	return acc
}
