/*
	https://leetcode.com/problems/short-encoding-of-words/

	A valid encoding of an array of words is any reference string s and array of indices
	indices such that:

		words.length == indices.length
		The reference string s ends with the '#' character.
		For each index indices[i], the substring of s starting from indices[i] and up to
	(but not including) the next '#' character is equal to words[i].

	Given an array of words, return the length of the shortest reference string s possible
	of any valid encoding of words
*/

package solutions

func minimumLengthEncoding(words []string) int {
	good := make(Set[string])
	for _, w := range words {
		good.Add(w)
	}
	for w := range good {
		for k := 1; k < len(w); k++ {
			good.Remove(w[k:])
		}
	}

	ans := 0
	for w := range good {
		ans += len(w) + 1
	}
	return ans
}
