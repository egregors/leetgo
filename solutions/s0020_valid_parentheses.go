/*
	https://leetcode.com/problems/valid-parentheses/

	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
	determine if the input string is valid.

	An input string is valid if:

		Open brackets must be closed by the same type of brackets.
		Open brackets must be closed in the correct order.
*/

package solutions

// S just a dummy stack
type S struct {
	stack []rune
}

func (s S) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *S) push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *S) pop() (rune, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	l := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return l, true
}

func isPaar(a, b rune) bool {
	switch a {
	case '(':
		return b == ')'
	case '{':
		return b == '}'
	case '[':
		return b == ']'
	default:
		return false
	}
}

func isL(r rune) bool {
	m := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}
	return m[r]
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := &S{}

	for _, r := range s {
		if isL(r) {
			stack.push(r)
			continue
		} else {
			if l, ok := stack.pop(); ok {
				if isPaar(l, r) {
					continue
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}

	return stack.isEmpty()
}
