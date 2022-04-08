/*
	https://leetcode.com/problems/kth-largest-element-in-a-stream/

	Design a class to find the kth largest element in a stream. Note that it is the kth
	largest element in the sorted order, not the kth distinct element.

	Implement KthLargest class:

		KthLargest(int k, int[] nums) Initializes the object with the integer k and the
		stream of integers nums.
		int add(int val) Appends the integer val to the stream and returns the element
		representing the kth largest element in the stream.
*/

package solutions

import (
	"container/heap"
)

type klHeap []int

func (k klHeap) Len() int           { return len(k) }
func (k klHeap) Less(i, j int) bool { return k[i] < k[j] }
func (k klHeap) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k klHeap) PeekL() int         { return k[0] }
func (k *klHeap) Push(x any)        { *k = append(*k, x.(int)) }
func (k *klHeap) Pop() any {
	x := (*k)[k.Len()-1]
	*k = (*k)[:k.Len()-1]
	return x
}

//nolint:revive // it's ok
type KthLargest struct {
	k int
	h klHeap
}

// NewKthLargest should call `Constructor` to pass LeetCode tests
func NewKthLargest(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
		h: klHeap{},
	}

	for _, n := range nums {
		kl.Add(n)
	}

	return kl
}

//nolint:revive // it's ok
func (kl *KthLargest) Add(val int) int {
	if kl.h.Len() < kl.k {
		heap.Push(&kl.h, val)
	} else if val > kl.h.PeekL() {
		heap.Pop(&kl.h)
		heap.Push(&kl.h, val)
	}
	return kl.h.PeekL()
}
