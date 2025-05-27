/*
	https://leetcode.com/problems/furthest-building-you-can-reach/

	You are given an integer array heights representing the heights of buildings,
		some bricks, and some ladders.

	You start your journey from building 0 and move to the next building by
		possibly using bricks or ladders.

	While moving from building i to building i+1 (0-indexed),

		If the current building's height is greater than or equal to the next
			building's height, you do not need a
			ladder or bricks.
		If the current building's height is less than the next building's height, you
			can either use one ladder or
			(h[i+1] - h[i]) bricks.

	Return the furthest building index (0-indexed) you can reach if you use the
		given ladders and bricks optimally.
*/

//nolint:revive // it's ok
package solutions

import "container/heap"

type RevIntHeap []int

func (h RevIntHeap) Len() int           { return len(h) }
func (h RevIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h RevIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds element into heap
func (h *RevIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// Pop removes the most small element of the heap from the heap end returns this element
func (h *RevIntHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func furthestBuilding(heights []int, bricks, ladders int) int {

	ih := &RevIntHeap{}
	for i := 0; i < len(heights)-1; i++ {
		gap := heights[i+1] - heights[i]
		if gap <= 0 {
			continue
		}
		heap.Push(ih, gap)
		bricks -= gap
		if bricks < 0 && ladders == 0 {
			return i
		}

		if bricks < 0 {
			bricks += heap.Pop(ih).(int)
			ladders--
		}
	}

	return len(heights) - 1
}
