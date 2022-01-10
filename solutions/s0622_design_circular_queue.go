/*
	https://leetcode.com/problems/design-circular-queue/

	Design your implementation of the circular queue. The circular queue is a linear data structure in which
	the operations are performed based on FIFO (First In First Out) principle and the last position is connected
	back to the first position to make a circle. It is also called "Ring Buffer".

	One of the benefits of the circular queue is that we can make use of the spaces in front of the queue.
	In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space
	in front of the queue. But using the circular queue, we can use the space to store new values.

	Implementation the MyCircularQueue class:

		MyCircularQueue(k) Initializes the object with the size of the queue to be k.
		int Front() Gets the front item from the queue. If the queue is empty, return -1.
		int Rear() Gets the last item from the queue. If the queue is empty, return -1.
		boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
		boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
		boolean isEmpty() Checks whether the circular queue is empty or not.
		boolean isFull() Checks whether the circular queue is full or not.

	You must solve the problem without using the built-in queue data structure in your programming language.
*/

// nolint:revive,stylecheck // das ist ok
package solutions

import (
	"sync"
)

type MyCircularQueue struct {
	q        []int
	len, cap int
	head     int

	mu sync.Mutex
}

// NewMyCircularQueue should call "Constructor" to pass LeetCode tests
func NewMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		q:    make([]int, k),
		len:  0,
		cap:  k,
		head: 0,
		mu:   sync.Mutex{},
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.IsFull() {
		return false
	}

	tailID := (q.head + q.len) % q.cap

	q.q[tailID] = value
	q.len++

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.IsEmpty() {
		return false
	}

	q.q[q.head] = -1
	q.head = (q.head + 1) % q.cap
	q.len--

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.q[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.q[(q.head+q.len-1)%q.cap]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.len == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.len == q.cap
}
