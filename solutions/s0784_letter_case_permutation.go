/*
	https://leetcode.com/problems/letter-case-permutation/

	Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.

	Return a list of all possible strings we could create. Return the output in any order.
*/

package solutions

import (
	"strings"
	"unicode"
)

func rec(xs string) []string {
	var out []string
	if xs != "" {
		if unicode.IsLetter(rune(xs[0])) {
			for _, leaf := range rec(xs[1:]) {
				out = append(
					out,
					strings.ToUpper(string(xs[0]))+leaf,
					strings.ToLower(string(xs[0]))+leaf,
				)
			}
		} else {
			for _, leaf := range rec(xs[1:]) {
				out = append(out, string(xs[0])+leaf)
			}
		}
	} else {
		out = append(out, "")
	}
	return out
}

func letterCasePermutation(s string) []string {
	return rec(s)
}
