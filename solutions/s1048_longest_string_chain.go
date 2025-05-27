/*
	https://leetcode.com/problems/longest-string-chain/

	You are given an array of words where each word consists of lowercase English
		letters.

	wordA is a predecessor of wordB if and only if we can insert exactly one letter
		anywhere
	in wordA without changing the order of the other characters to make it equal to
		wordB.

		For example, "abc" is a predecessor of "abac", while "cba" is not a
			predecessor of "bcad".

	A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1,
		where word1
	is a predecessor of word2, word2 is a predecessor of word3, and so on. A single
		word is
	trivially a word chain with k == 1.

	Return the length of the longest possible word chain with words chosen from the
		given list
	of words.
*/

package solutions

import (
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dp := make(map[string]int)
	res := 1

	for _, w := range words {
		currLen := 1
		for i := 0; i < len(w); i++ {
			temp := w
			temp = temp[:i] + temp[i+1:]
			prevLen := dp[temp]
			currLen = Maximum(currLen, prevLen+1)
		}
		dp[w] = currLen
		res = Maximum(res, currLen)
	}

	return res
}
