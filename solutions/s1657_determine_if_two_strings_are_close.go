/*
	https://leetcode.com/problems/determine-if-two-strings-are-close/

	Two strings are considered close if you can attain one from the other using the following
	operations:

		Operation 1: Swap any two existing characters.
			For example, abcde -> aecdb
		Operation 2: Transform every occurrence of one existing character into another existing
			character, and do the same with the other character.
			For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)

	You can use the operations on either string as many times as necessary.

	Given two strings, word1 and word2, return true if word1 and word2 are close, and false
	otherwise.
*/

package solutions

import "sort"

func closeStrings(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Map := make([]int, 26)
	word2Map := make([]int, 26)
	for _, r := range word1 {
		word1Map[r-'a']++
	}
	for _, r := range word2 {
		word2Map[r-'a']++
	}
	for i := 0; i < 26; i++ {
		if word1Map[i] == 0 && word2Map[i] > 0 ||
			word2Map[i] == 0 && word1Map[i] > 0 {
			return false
		}
	}
	sort.Ints(word1Map)
	sort.Ints(word2Map)

	for i := 0; i < 26; i++ {
		if word1Map[i] != word2Map[i] {
			return false
		}
	}
	return true
}
