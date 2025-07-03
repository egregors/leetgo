/*
	https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/

	Alice and Bob are playing a game. Initially, Alice has a string word = "a".

	You are given a positive integer k.

	Now Bob will ask Alice to perform the following operation forever:

		Generate a new string by changing each character in word to its next character
			in the English alphabet, and append it to the original word.

	For example, performing the operation on "c" generates "cd" and performing the
		operation on "zb" generates "zbac".

	Return the value of the kth character in word, after enough operations have
		been done for word to have at least k characters.

	Note that the character 'z' can be changed to 'a' in the operation.
*/

package solutions

import (
	"strings"
)

var s string

func kthCharacter(k int) byte {
	if s == "" {
		sb := strings.Builder{}
		sb.WriteRune('a')
		for sb.Len() < 500 {
			for _, ch := range sb.String() {
				sb.WriteRune(ch + 1%26)
			}
		}
		s = sb.String()
	}

	return s[k-1]
}
