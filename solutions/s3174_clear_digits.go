/*
	https://leetcode.com/problems/clear-digits/

	You are given a string s.

	Your task is to remove all digits by doing this operation repeatedly:

		Delete the first digit and the closest non-digit character to its left.

	Return the resulting string after removing all digits.

	Note that the operation cannot be performed on a digit that does not have any
		non-digit character to its left.
*/

package solutions

import "unicode"

func clearDigits(s string) string {
	st := []rune{rune(s[0])}

	for i := 1; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			st = append(st, rune(s[i]))
		} else {
			st = st[:len(st)-1]
		}
	}

	return string(st)
}
