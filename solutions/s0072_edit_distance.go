/*
	https://leetcode.com/problems/edit-distance/

	Given two strings word1 and word2, return the minimum number of operations
		required to convert word1 to word2.

	You have the following three operations permitted on a word:

		Insert a character
		Delete a character
		Replace a character

*/

package solutions

// minDistance72 should call minDistance to pass LeetCode tests
func minDistance72(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	var f func(string, string, int, int) int
	f = func(w1, w2 string, i, j int) int {
		if i == 0 {
			return j
		}
		if j == 0 {
			return i
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}

		midDist := 0
		if w1[i-1] == w2[j-1] {
			midDist = f(w1, w2, i-1, j-1)
		} else {
			insertDist := f(w1, w2, i, j-1)
			deleteDist := f(w1, w2, i-1, j)
			editDist := f(w1, w2, i-1, j-1)
			midDist = Minimum(insertDist, deleteDist, editDist) + 1
		}
		dp[i][j] = midDist
		return midDist
	}

	return f(word1, word2, len(word1), len(word2))
}
