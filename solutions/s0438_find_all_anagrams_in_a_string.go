/*
	https://leetcode.com/problems/find-all-anagrams-in-a-string/

	Given two strings s and p, return an array of all the start indices of p's anagrams in s.
	You may return the answer in any order.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
	typically using all the original letters exactly once.
*/

package solutions

func findAnagrams(s, p string) []int {
	m := make(map[[26]byte]int)

	pM := [26]byte{}
	for _, ch := range p {
		pM[ch-'a']++
	}
	m[pM]++

	var res []int
	lo, hi := 0, len(p)
	for ; hi <= len(s); lo, hi = lo+1, hi+1 {
		buf := [26]byte{}
		for _, ch := range s[lo:hi] {
			buf[ch-'a']++
		}
		if _, ok := m[buf]; ok {
			res = append(res, lo)
		}

	}

	return res
}
