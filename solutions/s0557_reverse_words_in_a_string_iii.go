/*
	https://leetcode.com/problems/reverse-words-in-a-string-iii/

	Given a string s, reverse the order of characters in each word within a
	sentence while still preserving whitespace and initial word order.
*/

package solutions

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	res := []string{}
	for _, w := range words {
		res = append(res, revWord([]byte(w)))
	}
	return strings.Join(res, " ")
}

func revWord(w []byte) string {
	lo, hi := 0, len(w)-1
	for lo <= hi {
		w[lo], w[hi] = w[hi], w[lo]
		lo++
		hi--
	}
	return string(w)
}
