/*
	https://leetcode.com/problems/pascals-triangle/

	Given an integer numRows, return the first numRows of Pascal's triangle.

	In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
	https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif
*/

package solutions

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		res[i] = append(res[i], 1)
		for j := 1; j < i; j++ {
			res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
		}
		res[i] = append(res[i], 1)
	}

	return res
}
