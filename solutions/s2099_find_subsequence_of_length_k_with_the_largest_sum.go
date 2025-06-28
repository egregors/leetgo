/*
	https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum

	You are given an integer array nums and an integer k. You want to find a
		subsequence of nums of length k that has the largest sum.

	Return any such subsequence as an integer array of length k.

	A subsequence is an array that can be derived from another array by deleting
		some or no elements without changing the order of the remaining elements.
*/

package solutions

import (
	"container/heap"
)

func maxSubsequence(nums []int, k int) []int {
	h := IntMinHeap{}

	for _, n := range nums {
		heap.Push(&h, n)
		for h.Len() > k {
			heap.Pop(&h)
		}
	}

	res := make([]int, 0, h.Len())
	m := make(map[int]int)
	for h.Len() > 0 {
		m[heap.Pop(&h).(int)]++
	}

	for _, n := range nums {
		if c, ok := m[n]; ok && c > 0 {
			res = append(res, n)
			m[n]--
		}
	}

	return res
}
