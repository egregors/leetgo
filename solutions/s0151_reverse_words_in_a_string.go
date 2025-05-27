/*
	https://leetcode.com/problems/reverse-words-in-a-string/

	Given an input string s, reverse the order of the words.

	A word is defined as a sequence of non-space characters. The words in s will be
	separated by at least one space.

	Return a string of the words in reverse order concatenated by a single space.

	Note that s may contain leading or trailing spaces or multiple spaces between
		two words.
		The returned string should only have a single space separating the words. Do
			not include any extra spaces.
*/

package solutions

import "strings"

func reverseWords151(s string) string {
	rev := func(ws []string) (res []string) {
		for i := len(ws) - 1; i >= 0; i-- {
			res = append(res, ws[i])
		}
		return res
	}

	filter := func(ws []string) (res []string) {
		for _, w := range ws {
			if w != "" {
				res = append(res, w)
			}
		}
		return res
	}
	return strings.Join(filter(rev(strings.Split(s, " "))), " ")
}
