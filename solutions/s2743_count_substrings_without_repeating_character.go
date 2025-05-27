/*
	https://leetcode.com/problems/count-substrings-without-repeating-character/description

	You are given a string s consisting only of lowercase English letters. We call
		a substring special if it
	contains no character which has occurred at least twice (in other words, it
		does not contain a repeating character).
	Your task is to count the number of special substrings. For example, in the
		string "pop", the substring "po" is a
	special substring, however, "pop" is not special (since 'p' has occurred
		twice).

	Return the number of special substrings.

	A substring is a contiguous sequence of characters within a string. For
		example, "abc" is a substring of "abcd", but
		"acd" is not.
*/

package solutions

func numberOfSpecialSubstrings(s string) int {
	c := 0
	freq := make([]int, 26)
	l, r := 0, 0
	for ; r < len(s); r++ {
		freq[s[r]-'a']++
		for freq[s[r]-'a'] > 1 {
			freq[s[l]-'a']--
			l++
		}
		c += r - l + 1
	}

	return c
}
