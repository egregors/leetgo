/*
	https://leetcode.com/problems/seat-reservation-manager/description/

	Design a system that manages the reservation state of n seats that are numbered from 1 to n.

	Implement the SeatManager class:

		SeatManager(int n) Initializes a SeatManager object that will manage n seats numbered from 1 to n.
	All seats are initially available.
		int reserve() Fetches the smallest-numbered unreserved seat, reserves it, and returns its number.
		void unreserve(int seatNumber) Unreserves the seat with the given seatNumber.

*/

//nolint:revive // it's ok
package solutions

import "container/heap"

type SeatManager struct {
	h *IntMinHeap
}

// NewSeatManager should call Constructor to pass Leetcode test
func NewSeatManager(n int) SeatManager {
	sm := SeatManager{
		h: &IntMinHeap{},
	}

	for i := 1; i <= n; i++ {
		heap.Push(sm.h, i)
	}

	return sm
}

func (sm *SeatManager) Reserve() int {
	return heap.Pop(sm.h).(int)
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	heap.Push(sm.h, seatNumber)
}
