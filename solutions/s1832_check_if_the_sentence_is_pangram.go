/*
	https://leetcode.com/problems/check-if-the-sentence-is-pangram/

	A pangram is a sentence where every letter of the English alphabet appears at
		least once.

	Given a string sentence containing only lowercase English letters, return true
		if sentence
	is a pangram, or false otherwise.
*/

package solutions

func checkIfPangram(sentence string) bool {
	ls := make([]rune, 26)
	for _, l := range sentence {
		ls[l-'a']++
	}
	for _, l := range ls {
		if l == 0 {
			return false
		}
	}
	return true
}
