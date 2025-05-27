/*
	https://leetcode.com/problems/backspace-string-compare/

	Given two strings s and t, return true if they are equal when both are typed
		into empty text editors.
	'#' means a backspace character.

	Note that after backspacing an empty text, the text will continue empty.
*/

package solutions

func buildStr(s string) string {
	var res []rune
	for _, ch := range s {
		if ch == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func backspaceCompare(s, t string) bool {
	return buildStr(s) == buildStr(t)
}
