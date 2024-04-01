/*
	https://leetcode.com/problems/length-of-last-word/

	Given a string s consisting of some words separated by some number of spaces,
	return the length of the last word in the string.

	A word is a maximal substring consisting of non-space characters only.
*/

package solutions

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' { res++ } else {
			if res == 0 { continue }
			return res
		}
	}

	return res
}
