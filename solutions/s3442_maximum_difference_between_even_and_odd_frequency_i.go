/*
	https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/

	You are given a string s consisting of lowercase English letters.

	Your task is to find the maximum difference diff = freq(a1) - freq(a2) between
		the frequency of characters a1 and a2 in the string such that:

		a1 has an odd frequency in the string.
		a2 has an even frequency in the string.

	Return this maximum difference.
*/

package solutions

import (
	"math"
)

func maxDifference(s string) int {
	m := make([]int, 26)
	for _, ch := range s {
		m[ch-'a']++
	}
	oddMax := 0
	evenMin := math.MaxInt

	for _, fr := range m {
		if fr != 0 {
			if fr%2 != 0 {
				oddMax = max(oddMax, fr)
			} else {
				evenMin = min(evenMin, fr)
			}
		}
	}

	return oddMax - evenMin
}
