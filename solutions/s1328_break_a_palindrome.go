/*
	https://leetcode.com/problems/break-a-palindrome/

	Given a palindromic string of lowercase English letters palindrome, replace exactly one character with any
	lowercase English letter so that the resulting string is not a palindrome and that it is the lexicographically
	smallest one possible.

	Return the resulting string. If there is no way to replace a character to make it not a palindrome, return an
	empty string.

	A string a is lexicographically smaller than a string b (of the same length) if in the first position where a
	and b differ, a has a character strictly smaller than the corresponding character in b. For example, "abcc" is
	lexicographically smaller than "abcd" because the first position they differ is at the fourth character, and 'c'
	is smaller than 'd'.
*/

package solutions

func breakPalindrome(palindrome string) string {
	if len(palindrome) < 2 {
		return ""
	}
	res := []rune(palindrome)

	for i := 0; i < len(res)/2; i++ {
		if res[i] != 'a' {
			res[i] = 'a'
			return string(res)
		}
	}

	res[len(res)-1] = 'b'
	return string(res)
}
