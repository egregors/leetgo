/*
	https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters

	You are given a string s consisting of lowercase English letters, and an
		integer k.

	Your task is to delete some (possibly none) of the characters in the string so
		that the number of distinct characters in the resulting string is at most k.

	Return the minimum number of deletions required to achieve this.
*/

package solutions

import "sort"

func minDeletion(s string, k int) int {
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}

	if len(m) <= k {
		return 0
	}

	xs := make([]int, 0, 8)
	for _, v := range m {
		xs = append(xs, v)
	}

	sort.Ints(xs)

	cnt := 0

	for i := 0; i < len(xs)-k; i++ {
		cnt += xs[i]
	}

	return cnt
}
