/*
	https://leetcode.com/problems/min-stack/

	Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

	Implement the MinStack class:

		MinStack() initializes the stack object.
		void push(int val) pushes the element val onto the stack.
		void pop() removes the element on the top of the stack.
		int top() gets the top element of the stack.
		int getMin() retrieves the minimum element in the stack.
*/

//nolint:revive // meh
package solutions

import "container/heap"

type MinStack struct {
	xs   []int
	mins *IntMinHeap
}

// NewMinStack should call Constructor to pass LeetCode tests
func NewMinStack() MinStack {
	return MinStack{xs: make([]int, 0), mins: &IntMinHeap{}}
}

func (s *MinStack) Push(val int) {
	s.xs = append(s.xs, val)
	heap.Push(s.mins, val)
}

func (s *MinStack) Pop() {
	s.xs = s.xs[:len(s.xs)-1]
}

func (s *MinStack) Top() int {
	return s.xs[len(s.xs)-1]
}

func (s *MinStack) GetMin() int {
	// TODO: i need to make this method O(1)...
	return Minimum(s.xs...)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
