/*
	https://leetcode.com/problems/length-of-last-word/

	Given a string s consisting of some words separated by some number of spaces,
	return the length of the last word in the string.

	A word is a maximal substring consisting of non-space characters only.
*/

package solutions

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 {
		if s[i] != ' ' {
			j := i
			for j >= 0 && s[j] != ' ' {
				j--
			}
			return i - j
		}
		i--
	}

	return 0
}
