/*
	https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

	Given a string s of '(' , ')' and lowercase English characters.

	Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the
	resulting parentheses string is valid and return any valid string.

	Formally, a parentheses string is valid if and only if:

		It is the empty string, contains only lowercase characters, or
		It can be written as AB (A concatenated with B), where A and B are valid strings, or
		It can be written as (A), where A is a valid string.
*/

package solutions

func del(xs *[]byte, i int) {
	*xs = append((*xs)[:i], (*xs)[i+1:]...)
}

func minRemoveToMakeValid(s string) string {
	xs, open, i := []byte(s), 0, 0

	for i < len(xs) {
		if xs[i] == '(' {
			open++
		} else if xs[i] == ')' {
			if open == 0 {
				del(&xs, i)
				i--
			} else {
				open--
			}
		}
		i++
	}

	i = len(xs) - 1
	for open > 0 && i >= 0 {
		if xs[i] == '(' {
			del(&xs, i)
			open--
		}
		i--
	}

	return string(xs)
}
