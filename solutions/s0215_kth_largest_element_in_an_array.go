/*
	https://leetcode.com/problems/kth-largest-element-in-an-array/

	Given an integer array nums and an integer k, return the kth largest element in the array.

	Note that it is the kth largest element in the sorted order, not the kth distinct element.
*/

//nolint:revive // it's ok
package solutions

import "container/heap"

func findKthLargest(nums []int, k int) int {
	ih := &IntMinHeap{}
	for _, n := range nums {
		heap.Push(ih, n)
		if ih.Len() > k {
			heap.Pop(ih)
		}
	}
	return heap.Pop(ih).(int)
}
