/*
	https://leetcode.com/problems/word-break/description/

	Given a string s and a dictionary of strings wordDict, return true if s can be
	segmented into a space-separated sequence of one or more dictionary words.

	Note that the same word in the dictionary may be reused multiple times in the
		segmentation.
*/

package solutions

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	words := NewSet(wordDict)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words.Contains(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
