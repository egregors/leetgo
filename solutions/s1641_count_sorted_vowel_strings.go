/*
	https://leetcode.com/problems/count-sorted-vowel-strings/

	Given an integer n, return the number of strings of length n that consist only
		of
	vowels (a, e, i, o, u) and are lexicographically sorted.

	A string s is lexicographically sorted if for all valid i, s[i] is the same as
		or comes
	before s[i+1] in the alphabet.
*/

package solutions

func countVowelStrings(n int) int {
	// https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
