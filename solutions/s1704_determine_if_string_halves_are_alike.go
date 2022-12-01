/*
	https://leetcode.com/problems/determine-if-string-halves-are-alike/

	You are given a string s of even length. Split this string into two halves of equal
	lengths, and let a be the first half and b be the second half.

	Two strings are alike if they have the same number of vowels
	('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s
	contains uppercase and lowercase letters.

	Return true if a and b are alike. Otherwise, return false.
*/

package solutions

func halvesAreAlike(s string) bool {
	lowels := NewSet([]byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'})
	n := len(s) / 2
	l, r := s[:n], s[n:]
	var cl, cr int
	for i := 0; i < n; i++ {
		if lowels.Contains(l[i]) {
			cl++
		}
		if lowels.Contains(r[i]) {
			cr++
		}
	}

	return cl == cr
}
