/*
	https://leetcode.com/problems/delete-operation-for-two-strings/

	Given two strings word1 and word2, return the minimum number of steps required
		to make word1 and word2 the same.

	In one step, you can delete exactly one character in either string.
*/

package solutions

func minDistance(word1, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := 0; i <= len(word1); i++ {
		temp := make([]int, len(word2)+1)
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				temp[j] = i + j
			} else if word1[i-1] == word2[j-1] {
				temp[j] = dp[j-1]
			} else {
				temp[j] = 1 + Minimum(dp[j], temp[j-1])
			}
		}
		dp = temp
	}
	return dp[len(word2)]
}
