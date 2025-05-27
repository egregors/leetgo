/*
	https://leetcode.com/problems/implement-stack-using-queues/

	Implement a last-in-first-out (LIFO) stack using only two queues. The
		implemented stack
	should support all the functions of a normal stack (push, top, pop, and empty).

	Implement the MyStack class:

		void push(int x) Pushes element x to the top of the stack.
		int pop() Removes the element on the top of the stack and returns it.
		int top() Returns the element on the top of the stack.
		boolean empty() Returns true if the stack is empty, false otherwise.

	Notes:

		You must use only standard operations of a queue, which means that only push
			to back,
		peek/pop from front, size and is empty operations are valid.
		Depending on your language, the queue may not be supported natively. You may
			simulate a
		queue using a list or deque (double-ended queue) as long as you use only a
			queue's
		standard operations.
*/

//nolint:revive // it's ok
package solutions

type MyQueue225 []int

func (q *MyQueue225) Push(x int)    { *q = append(*q, x) }
func (q *MyQueue225) Size() int     { return len(*q) }
func (q *MyQueue225) IsEmpty() bool { return q.Size() == 0 }
func (q *MyQueue225) Peek() int {
	if !q.IsEmpty() {
		return (*q)[0]
	}
	return -1
}
func (q *MyQueue225) Pop() int {
	x := q.Peek()
	*q = (*q)[1:]
	return x
}

type MyStack struct {
	q MyQueue225
}

// NewMyStack should call Constructor to pass LeetCode tests
func NewMyStack() MyStack {
	return MyStack{MyQueue225{}}
}

func (s *MyStack) Push(x int) {
	s.q.Push(x)
	sz := s.q.Size()
	for sz > 1 {
		s.q.Push(s.q.Pop())
		sz--
	}
}

func (s *MyStack) Pop() int {
	return s.q.Pop()
}

func (s *MyStack) Top() int {
	return s.q.Peek()
}

func (s *MyStack) Empty() bool {
	return s.q.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
