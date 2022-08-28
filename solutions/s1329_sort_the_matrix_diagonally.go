/*
	https://leetcode.com/problems/sort-the-matrix-diagonally/

	A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost
	row or leftmost column and going in the bottom-right direction until reaching the matrix's end.
	For example, the matrix diagonal starting from mat[2][0], where mat is a 6 x 3 matrix,
	includes cells mat[2][0], mat[3][1], and mat[4][2].

	Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return the
	resulting matrix.
*/

package solutions

import (
	"container/heap"
)

func diagonalSort(mat [][]int) [][]int {
	m := make(map[int]*IntMinHeap)

	for r := 0; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if _, ok := m[r-c]; !ok {
				m[r-c] = &IntMinHeap{}
			}
			heap.Push(m[r-c], mat[r][c])
		}
	}

	for r := 0; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			mat[r][c] = heap.Pop(m[r-c]).(int)
		}
	}

	return mat
}
