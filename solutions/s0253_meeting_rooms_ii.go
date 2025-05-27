/*
	https://leetcode.com/problems/meeting-rooms-ii/description

	Given an array of meeting time intervals intervals where intervals[i] =
		[starti, endi], return the minimum
	number of conference rooms required.
*/

package solutions

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	maxRooms := 0
	h := &MinHeap253{}

	for _, i := range intervals {
		if h.Len() == 0 {
			// 0 free rooms
			heap.Push(h, i)
		} else {
			nextRoom := heap.Pop(h).([]int)
			if nextRoom[1] <= i[0] {
				// reuse free room
				heap.Push(h, i)
			} else {
				// call for another one
				heap.Push(h, nextRoom)
				heap.Push(h, i)
			}
		}
		maxRooms = max(maxRooms, h.Len())
	}

	return maxRooms
}

type MinHeap253 [][]int

func (h *MinHeap253) Len() int           { return len(*h) }
func (h *MinHeap253) Less(i, j int) bool { return (*h)[i][1] < (*h)[j][1] }
func (h *MinHeap253) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap253) Push(x any)         { *h = append(*h, x.([]int)) }
func (h *MinHeap253) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}
