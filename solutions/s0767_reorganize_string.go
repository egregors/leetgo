/*
	https://leetcode.com/problems/reorganize-string/

	Given a string s, rearrange the characters of s so that any two adjacent
	characters are not the same.

	Return any possible rearrangement of s or return "" if not possible.
*/

//nolint:revive // it's ok my friend
package solutions

import "container/heap"

type Pair767 struct {
	ch byte
	c  int
}

type MaxPairHeap767 []*Pair767

func (h *MaxPairHeap767) Len() int           { return len(*h) }
func (h *MaxPairHeap767) Less(i, j int) bool { return (*h)[i].c > (*h)[j].c }
func (h *MaxPairHeap767) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxPairHeap767) Push(x interface{}) { *h = append(*h, x.(*Pair767)) }
func (h *MaxPairHeap767) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func reorganizeString(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	h := &MaxPairHeap767{}
	for k, v := range m {
		heap.Push(h, &Pair767{k, v})
	}

	var res []byte
	var prev *Pair767
	for h.Len() > 0 {
		ch := heap.Pop(h).(*Pair767)
		if ch == prev {
			if h.Len() == 0 {
				return ""
			}
			next := heap.Pop(h).(*Pair767)
			heap.Push(h, ch)
			ch = next
		}
		res = append(res, ch.ch)
		ch.c--
		if ch.c != 0 {
			heap.Push(h, ch)
		}
		prev = ch
	}

	return string(res)
}
