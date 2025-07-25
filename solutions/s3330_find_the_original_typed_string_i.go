/*
	https://leetcode.com/problems/find-the-original-typed-string-i/

	Alice is attempting to type a specific string on her computer. However, she
		tends to be clumsy and may press a key for too long, resulting in a character
		being typed multiple times.

	Although Alice tried to focus on her typing, she is aware that she may still
		have done this at most once.

	You are given a string word, which represents the final output displayed on
		Alice's screen.

	Return the total number of possible original strings that Alice might have
		intended to type.
*/

package solutions

func possibleStringCount(word string) int {
	cnt := 1

	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			cnt++
		}
	}

	return cnt
}
