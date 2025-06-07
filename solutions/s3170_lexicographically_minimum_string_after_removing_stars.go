/*
	https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/

	You are given a string s. It may contain any number of '*' characters. Your
		task is to remove all '*' characters.

	While there is a '*', do the following operation:

		Delete the leftmost '*' and the smallest non-'*' character to its left. If
			there are several smallest characters, you can delete any of them.

	Return the LGS resulting string after removing all '*' characters.
*/

package solutions

import "strings"

func clearStars(s string) string {
	if !strings.ContainsAny(s, "*") {
		return s
	}
	deleted := make(map[int]struct{})
	m := make([][]int, 26)
	for i, ch := range s {
		if ch != '*' {
			m[ch-'a'] = append(m[ch-'a'], i)
		} else {
			for l, is := range m {
				if len(is) > 0 {
					idx := is[len(is)-1]
					deleted[idx] = struct{}{}
					m[l] = is[:len(is)-1]
					break
				}
			}
		}
	}

	sb := strings.Builder{}
	for i, ch := range s {
		if _, ok := deleted[i]; ok || ch == '*' {
			continue
		}
		sb.WriteRune(ch)
	}

	return sb.String()
}
