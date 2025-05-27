/*
	https://leetcode.com/problems/valid-parentheses/

	Given a string s containing just the characters '(', ')', '{', '}', '[' and
		']',
	determine if the input string is valid.

	An input string is valid if:

		Open brackets must be closed by the same type of brackets.
		Open brackets must be closed in the correct order.
*/

package solutions

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var xs []rune
	for _, c := range s {
		if b, ok := m[c]; ok {
			if len(xs) == 0 {
				return false
			}
			if xs[len(xs)-1] != b {
				return false
			}
			xs = xs[:len(xs)-1]
		} else {
			xs = append(xs, c)
		}
	}
	return len(xs) == 0
}
