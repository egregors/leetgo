/*
	https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/

	A string s is called good if there are no two different characters in s that
		have the same frequency.

	Given a string s, return the minimum number of characters you need to delete to
		make s good.

	The frequency of a character in a string is the number of times it appears in
		the string. For example,
	in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is
		1.
*/

package solutions

func minDeletions(s string) int {
	ls := make([]int, 26)
	for _, ch := range s {
		ls[ch-'a']++
	}

	m := make(map[int]struct{})
	del := 0

	for i := 0; i < 26; i++ {
		for _, ok := m[ls[i]]; ls[i] > 0 && ok; _, ok = m[ls[i]] {
			ls[i]--
			del++
		}
		m[ls[i]] = struct{}{}
	}

	return del
}
