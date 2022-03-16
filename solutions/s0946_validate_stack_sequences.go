/*
	https://leetcode.com/problems/validate-stack-sequences/

	Given two integer arrays pushed and popped each with distinct values,
	return true if this could have been the result of a sequence of push and
	pop operations on an initially empty stack946, or false otherwise.
*/

package solutions

type stack946 []int

func (s *stack946) IsEmpty() bool { return len(*s) == 0 }
func (s *stack946) Push(x int)    { *s = append(*s, x) }
func (s *stack946) Peek() int     { return (*s)[len(*s)-1] }
func (s *stack946) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func validateStackSequences(pushed, popped []int) bool {
	s, popID := stack946{}, 0
	for _, push := range pushed {
		if push != popped[popID] {
			s.Push(push)
		} else {
			popID++
			for !s.IsEmpty() && s.Peek() == popped[popID] {
				s.Pop()
				popID++
			}
		}
	}
	return s.IsEmpty()
}
