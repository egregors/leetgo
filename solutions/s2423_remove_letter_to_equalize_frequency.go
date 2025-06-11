/*
	https://leetcode.com/problems/remove-letter-to-equalize-frequency/

	You are given a 0-indexed string word, consisting of lowercase English letters.
		You need to select one index and remove the letter at that index from word so
		that the frequency of every letter present in word is equal.

	Return true if it is possible to remove one letter so that the frequency of all
		letters in word are equal, and false otherwise.

	Note:

		The frequency of a letter x is the number of times it occurs in the string.
		You must remove exactly one letter and cannot choose to do nothing.
*/

package solutions

func equalFrequency(word string) bool {
	m := make([]int, 26)
	for _, ch := range word {
		m[ch-'a']++
	}

	check := func(xs []int) bool {
		fr := 0
		for _, x := range xs {
			if x != 0 {
				if fr == 0 {
					fr = x
				}
				if fr != x {
					return false
				}
			}
		}

		return true
	}

	for i := 0; i < 26; i++ {
		if m[i] > 0 {
			m[i]--
			if check(m) {
				return true
			}
			m[i]++
		}
	}

	return false
}
