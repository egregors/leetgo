/*
	https://leetcode.com/problems/number-of-matching-subsequences/

	Given a string s and an array of strings words, return the number of words[i] that
	is a subsequence of s.

	A subsequence of a string is a new string generated from the original string with some
	characters (can be none) deleted without changing the relative order of the remaining
	characters.

		For example, "ace" is a subsequence of "abcde".
*/

package solutions

func numMatchingSubseq(s string, words []string) int {
	m := make(map[rune][]int)
	for i, ch := range s {
		m[ch] = append(m[ch], i)
	}
	count := 0

LOOP:
	for _, word := range words {
		if ids, ok := m[rune(word[0])]; ok {
			curr, found := ids[0], 1
			for i := 1; i < len(word); i++ {
				ch := rune(word[i])
				if ids, ok = m[ch]; ok {
					for _, idx := range ids {
						if idx > curr {
							curr = idx
							found++
							break
						}
					}
				} else {
					continue LOOP
				}
			}
			if found == len(word) {
				count++
			}
		} else {
			continue
		}
	}

	return count
}
