/*
	https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

	You are given a string s consisting of lowercase English letters ('a' to 'z').

	Your task is to:

		Find the vowel (one of 'a', 'e', 'i', 'o', or 'u') with the maximum frequency.
		Find the consonant (all other letters excluding vowels) with the maximum
			frequency.

	Return the sum of the two frequencies.

	Note: If multiple vowels or consonants have the same maximum frequency, you may
		choose any one of them. If there are no vowels or no consonants in the string,
		consider their frequency as 0.
	The frequency of a letter x is the number of times it occurs in the string.
*/

package solutions

func maxFreqSum(s string) int {
	m := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}

	var maxL, maxC int
	for k, v := range cnt {
		if m[k] {
			maxL = max(maxL, v)
			continue
		}
		maxC = max(maxC, v)
	}

	return maxL + maxC
}
