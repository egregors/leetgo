/*
	https://leetcode.com/problems/top-k-frequent-words/

	Given an array of strings words and an integer k, return the k most frequent
		strings.

	Return the answer sorted by the frequency from highest to lowest. Sort the
		words with the
	same frequency by their lexicographical order.
*/

//nolint:revive // it's ok
package solutions

import (
	"container/heap"
)

type Word692 struct {
	w string
	c int
}

type WordHeap692 []Word692

func (w *WordHeap692) Len() int { return len(*w) }
func (w *WordHeap692) Less(i, j int) bool {
	if (*w)[i].c == (*w)[j].c {
		return (*w)[i].w > (*w)[j].w
	}
	return (*w)[i].c < (*w)[j].c
}
func (w *WordHeap692) Swap(i, j int) { (*w)[i], (*w)[j] = (*w)[j], (*w)[i] }
func (w *WordHeap692) Push(x any)    { *w = append(*w, x.(Word692)) }
func (w *WordHeap692) Pop() any {
	x := (*w)[w.Len()-1]
	*w = (*w)[:w.Len()-1]
	return x
}

// topKFrequent692 should call topKFrequent to pass LeetCode tests
func topKFrequent692(words []string, k int) []string {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	h := &WordHeap692{}
	for word, count := range m {
		heap.Push(h, Word692{word, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(Word692).w
	}

	return res
}
