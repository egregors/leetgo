/*
	https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

	You are given a string s consisting of lowercase English letters. A duplicate
		removal consists of choosing
	two adjacent and equal letters and removing them.

	We repeatedly make duplicate removals on s until we no longer can.

	Return the final string after all such duplicate removals have been made. It
		can be proven that the answer is unique.

*/

//nolint:revive // it's ok
package solutions

type Stack1047 []rune

func (s *Stack1047) IsEmpty() bool   { return len(*s) == 0 }
func (s *Stack1047) Peek() rune      { return (*s)[len(*s)-1] }
func (s *Stack1047) Push(ch rune)    { *s = append(*s, ch) }
func (s *Stack1047) ToSlice() []rune { return *s }
func (s *Stack1047) Pop() rune {
	ch := s.Peek()
	*s = (*s)[:len(*s)-1]
	return ch
}

// removeDuplicates1047 could call removeDuplicates to pass Leetcode tests
func removeDuplicates1047(s string) string {
	stack := Stack1047{}
	for _, ch := range s {
		if !stack.IsEmpty() && stack.Peek() == ch {
			stack.Pop()
		} else {
			stack.Push(ch)
		}
	}
	return string(stack.ToSlice())
}
