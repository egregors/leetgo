/*
	https://leetcode.com/problems/repeated-dna-sequences/

		The DNA sequence is composed of a series of nucleotides abbreviated as 'A',
			'C',
		'G', and 'T'.

		For example, "ACGAATTCCG" is a DNA sequence.

	When studying DNA, it is useful to identify repeated sequences within the DNA.
	Given a string s that represents a DNA sequence, return all the 10-letter-long
	sequences (substrings) that occur more than once in a DNA molecule. You may
		return
	the answer in any order.
*/

package solutions

func findRepeatedDnaSequences(s string) []string {
	xs := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		xs[s[i:i+10]]++
	}

	var res []string
	for k, v := range xs {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}
