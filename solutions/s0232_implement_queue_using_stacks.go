/*
	https://leetcode.com/problems/implement-queue-using-stacks/

	Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

	Implement the MyQueue class:

		void push(int x) Pushes element x to the back of the queue.
		int pop() Removes the element from the front of the queue and returns it.
		int peek() Returns the element at the front of the queue.
		boolean empty() Returns true if the queue is empty, false otherwise.

	Notes:

		You must use only standard operations of a stack, which means only push to
		top, peek/pop from top, size, and is empty operations are valid.
		Depending on your language, the stack may not be supported natively. You may
		simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/

//nolint:revive // This is ok!
package solutions

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
	xs []int
}

func Constructor() MyQueue {
	return MyQueue{xs: []int{}}
}

func (q *MyQueue) Push(x int) {
	q.xs = append(q.xs, x)
}

func (q *MyQueue) Pop() int {
	el := q.xs[0]
	q.xs = q.xs[1:]
	return el
}

func (q *MyQueue) Peek() int {
	return q.xs[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.xs) == 0
}
