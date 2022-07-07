/*
	https://leetcode.com/problems/interleaving-string/

	Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

	An interleaving of two strings s and t is a configuration where they are divided into
	non-empty substrings such that:

		s = s1 + s2 + ... + sn
		t = t1 + t2 + ... + tm
		|n - m| <= 1
		The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...

	Note: a + b is the concatenation of strings a and b.
*/

package solutions

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[j] = true
			} else if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s2)]
}
