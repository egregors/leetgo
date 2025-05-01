/*
	https://leetcode.com/problems/remove-all-occurrences-of-a-substring/description/

	Given two strings s and part, perform the following operation on s until all occurrences of the substring
	part are removed:

		Find the leftmost occurrence of the substring part and remove it from s.

	Return s after removing all occurrences of part.

	A substring is a contiguous sequence of characters in a string.
*/

package solutions

func removeOccurrences(s string, part string) string {
	st := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		st = append(st, s[i])
		if len(st) >= len(part) && string(st[len(st)-len(part):]) == part {
			st = st[:len(st)-len(part)]
		}
	}
	return string(st)
}
