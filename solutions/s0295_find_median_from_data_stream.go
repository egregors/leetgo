/*
	https://leetcode.com/problems/find-median-from-data-stream/

	The median is the middle value in an ordered integer list. If the size of the
		list is even, there
	is no middle value, and the median is the mean of the two middle values.

		For example, for arr = [2,3,4], the median is 3.
		For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

	Implement the MedianFinder class:

		MedianFinder() initializes the MedianFinder object.
		void addNum(int num) adds the integer num from the data stream to the data
			structure.
		double findMedian() returns the median of all elements so far. Answers within
			10-5 of the actual
	answer will be accepted.
*/

//nolint:revive // it's ok
package solutions

import (
	"container/heap"
)

type MedianFinder struct {
	lo IntMaxHeap
	hi IntMinHeap
}

// NewMedianFinder should call Constructor to pass Leetcode tests
func NewMedianFinder() MedianFinder {
	return MedianFinder{
		lo: IntMaxHeap{},
		hi: IntMinHeap{},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(&mf.lo, num)

	heap.Push(&mf.hi, mf.lo[0])
	heap.Pop(&mf.lo)

	if mf.lo.Len() < mf.hi.Len() {
		heap.Push(&mf.lo, mf.hi[0])
		heap.Pop(&mf.hi)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64(mf.lo[0])
	}
	return float64(mf.lo[0]+mf.hi[0]) * 0.5
}
