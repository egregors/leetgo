/*
	https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/?envType=daily-question&envId=2024-04-04

	A string is a valid parentheses string (denoted VPS) if it meets one of the following:

	It is an empty string "", or a single character not equal to "(" or ")",
	It can be written as AB (A concatenated with B), where A and B are VPS's, or
	It can be written as (A), where A is a VPS.

	We can similarly define the nesting depth depth(S) of any VPS S as follows:

			depth("") = 0
			depth(C) = 0, where C is a string with a single character not equal to "(" or ")".
			depth(A + B) = max(depth(A), depth(B)), where A and B are VPS's.
			depth("(" + A + ")") = 1 + depth(A), where A is a VPS.

	For example, "", "()()", and "()(()())" are VPS's (with nesting depths 0, 1, and 2), and ")(" and "(()" are not VPS's.

	Given a VPS represented as string s, return the nesting depth of s.
*/

package solutions

// maxDepth1614 shoult call maxDepth to pass leetcode test
func maxDepth1614(s string) int {
	depth := 0
	m := 0
	for _, ch := range s {
		if ch == '(' {
			depth++
			if depth > m {
				m = depth
			}
		}
		if ch == ')' {
			depth--
		}
	}

	return m
}
