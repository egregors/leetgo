/*
	https://leetcode.com/problems/construct-target-array-with-multiple-sums/

	You are given an array target of n integers. From a starting array arr consisting of n 1's,
	you may perform the following procedure :

		let x be the sum of all elements currently in your array.
		choose index i, such that 0 <= i < n and set the value of arr at index i to x.
		You may repeat this procedure as many times as needed.

	Return true if it is possible to construct the target array from arr, otherwise, return false.
*/

//nolint:revive // it's ok
package solutions

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func isPossible(target []int) bool {
	maxHeap, sum := new(MaxHeap), 0
	for _, n := range target {
		heap.Push(maxHeap, n)
		sum += n
	}

	for {
		a := heap.Pop(maxHeap).(int)
		sum -= a
		if a == 1 || sum == 1 {
			return true
		}
		if sum >= a || sum == 0 || a%sum == 0 {
			return false
		}
		a %= sum
		heap.Push(maxHeap, a)
		sum += a
	}
}
