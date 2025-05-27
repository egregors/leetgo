/*
	https://leetcode.com/problems/top-k-frequent-elements/

	Given an integer array nums and an integer k, return the k most frequent
		elements. You may return the answer in any order.
*/

package solutions

import "container/heap"

//nolint:revive // it's meh
type KVPair struct {
	Val, P int
}

type tHeap []*KVPair

func (h tHeap) Len() int           { return len(h) }
func (h tHeap) Less(i, j int) bool { return h[i].P < h[j].P }
func (h tHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tHeap) Push(x interface{}) {
	*h = append(*h, x.(*KVPair))
}

func (h *tHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	h := &tHeap{}
	for key, v := range m {
		heap.Push(h, &KVPair{Val: key, P: v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(*KVPair).Val)
	}

	return res
}
