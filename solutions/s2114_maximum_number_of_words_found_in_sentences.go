/*
	https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

	A sentence is a list of words that are separated by a single space with no
		leading or trailing spaces.

	You are given an array of strings sentences, where each sentences[i] represents
		a single sentence.

	Return the maximum number of words that appear in a single sentence.
*/

package solutions

import "strings"

func mostWordsFound(sentences []string) int {
	wc := 0
	for _, s := range sentences {
		if words := len(strings.Split(s, " ")); words > wc {
			wc = words
		}
	}
	return wc
}
