/*
	https://leetcode.com/problems/pascals-triangle-ii/

	Given an integer rowIndex, return the rowIndexth (0-indexed) row of the
		Pascal's triangle.
*/

package solutions

func getRow(rowIndex int) []int {
	res := []int{1}
	for rowIndex > 0 {
		res = append(res, 1)
		for i := len(res) - 1; i >= 2; i-- {
			res[i-1] = res[i-2] + res[i-1]
		}
		rowIndex--
	}
	return res
}
