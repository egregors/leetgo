/*
	https://leetcode.com/problems/rotate-image/

	You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

	You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
	DO NOT allocate another 2D matrix and do the rotation.
*/

package solutions

func rotateImage(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n>>1 + n%2); i++ {
		for j := 0; j < n>>1; j++ {
			tmp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}
