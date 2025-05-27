/*
	https://leetcode.com/problems/longest-common-subsequence/

	Given two strings text1 and text2, return the length of their longest common
		subsequence. If there is no common
	subsequence, return 0.

	A subsequence of a string is a new string generated from the original string
		with some characters (can be none)
	deleted without changing the relative order of the remaining characters.

		For example, "ace" is a subsequence of "abcde".

	A common subsequence of two strings is a subsequence that is common to both
		strings.
*/

package solutions

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for ri := range dp {
		dp[ri] = make([]int, m+1)
	}
	// dp[r][c] == 0 init state
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}
