/*
	https://leetcode.com/problems/find-words-that-can-be-formed-by-characters

	You are given an array of strings words and a string chars.

	A string is good if it can be formed by characters from chars (each character
		can only
	be used once for each word in words).

	Return the sum of lengths of all good strings in words.
*/

package solutions

func countCharacters(words []string, chars string) int {
	m := make([]int, 26)
	for _, ch := range chars {
		m[ch-'a']++
	}

	cnt := 0
LOOP:
	for _, w := range words {
		wM := make([]int, 26)
		for _, ch := range w {
			wM[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if m[i]-wM[i] < 0 {
				continue LOOP
			}
		}
		cnt += len(w)
	}

	return cnt
}
