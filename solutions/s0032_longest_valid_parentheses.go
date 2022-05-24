/*
	https://leetcode.com/problems/longest-valid-parentheses/

	Given a string containing just the characters '(' and ')',
	find the length of the longest valid (well-formed) parentheses substring.
*/

//nolint:revive // it's ok
package solutions

type Stack32 []int

func (s *Stack32) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack32) Push(x int)    { *s = append(*s, x) }
func (s *Stack32) Peek() int     { return (*s)[len(*s)-1] }
func (s *Stack32) Pop() int {
	x := s.Peek()
	*s = (*s)[:len(*s)-1]
	return x
}

func longestValidParentheses(s string) int {
	max := 0
	stack := Stack32{}
	stack.Push(-1)

	for i, ch := range s {
		if ch == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				max = Maximum(max, i-stack.Peek())
			}
		}
	}
	return max
}
