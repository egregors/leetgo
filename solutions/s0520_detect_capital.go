/*
	https://leetcode.com/problems/detect-capital/description/

		We define the usage of capitals in a word to be right when one of the
	following cases holds:

		All letters in this word are capitals, like "USA".
		All letters in this word are not capitals, like "leetcode".
		Only the first letter in this word is capital, like "Google".

	Given a string word, return true if the usage of capitals in it is right.
*/

package solutions

import "strings"

func detectCapitalUse(word string) bool {
	lower := strings.ToLower(word)
	cases := map[string]bool{
		strings.ToUpper(word): true,
		lower:                 true,
		strings.ToUpper(string(word[0])) + lower[1:]: true,
	}
	return cases[word]
}
