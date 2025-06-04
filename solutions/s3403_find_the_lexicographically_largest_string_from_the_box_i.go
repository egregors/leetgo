/*
	https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i

	You are given a string word, and an integer numFriends.

	Alice is organizing a game for her numFriends friends. There are multiple
		rounds in the game, where in each round:

		word is split into numFriends non-empty strings, such that no previous round
			has had the exact same split.
		All the split words are put into a box.

	Find the
	string from the box after all the rounds are finished.
*/

package solutions

func lastSubstring(s string) string {
	i, j, n := 0, 1, len(s)
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j = j + k + 1
		}
	}
	return s[i:]
}

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	last := lastSubstring(word)
	n, m := len(word), len(last)
	return last[:min(m, n-numFriends+1)]
}
