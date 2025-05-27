/*
	https://leetcode.com/problems/sort-characters-by-frequency/

	Given a string s, sort it in decreasing order based on the frequency of the
		characters.
	The frequency of a character is the number of times it appears in the string.

	Return the sorted string. If there are multiple answers, return any of them.
*/

package solutions

import "container/heap"

//nolint:revive // it's ok
type Char struct {
	Val rune
	P   int
}

type charHeap []*Char

func (h charHeap) Len() int           { return len(h) }
func (h charHeap) Less(i, j int) bool { return h[i].P > h[j].P }
func (h charHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *charHeap) Push(x interface{}) { *h = append(*h, x.(*Char)) }
func (h *charHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}

	h := &charHeap{}
	for k, v := range m {
		heap.Push(h, &Char{Val: k, P: v})
	}

	var res []rune
	for h.Len() > 0 {
		ch := heap.Pop(h).(*Char)
		var buf []rune
		for i := 0; i < ch.P; i++ {
			buf = append(buf, ch.Val)
		}
		res = append(res, buf...)
	}

	return string(res)
}
