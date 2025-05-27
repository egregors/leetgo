/*
	https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

	Given an n x n matrix where each of the rows and columns is sorted in ascending
		order,
	return the kth smallest element in the matrix.

	Note that it is the kth smallest element in the sorted order, not the kth
		distinct element.

	You must find a solution with a memory complexity better than O(n2).
*/

package solutions

import (
	"container/heap"
)

func kthSmallest378(matrix [][]int, k int) int {
	n := len(matrix)
	ih := &IntMaxHeap{}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			heap.Push(ih, matrix[r][c])
			if ih.Len() > k {
				heap.Pop(ih)
			}
		}
	}
	return heap.Pop(ih).(int)
}
