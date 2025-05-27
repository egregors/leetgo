/*
	https://leetcode.com/problems/repeated-substring-pattern/description/

	Given a string s, check if it can be constructed by taking a substring of it
		and appending
	multiple copies of the substring together.
*/

package solutions

import "strings"

func repeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}
