/*
	https://leetcode.com/problems/total-characters-in-string-after-transformations-i

	You are given a string s and an integer t, representing the number of transformations to perform. In one
	transformation, every character in s is replaced according to the following rules:

		If the character is 'z', replace it with the string "ab".
		Otherwise, replace it with the next character in the alphabet. For example, 'a' is replaced with 'b', 'b'
	is replaced with 'c', and so on.

	Return the length of the resulting string after exactly t transformations.

	Since the answer may be very large, return it modulo 109 + 7.
*/

package solutions

func lengthAfterTransformations(s string, t int) int {
	// | z -> ab
	// | a -> b -> c -> ... -> z
	MOD := 1_000_000_007
	cnt := 0

	dp := make([][]int, t+1)
	for ri := range len(dp) {
		dp[ri] = make([]int, 26)
	}
	for _, ch := range s {
		dp[0][ch-'a']++
	}

	for step := 1; step <= t; step++ {
		for i := range 26 {
			switch {
			case i < 25 && dp[step-1][i] > 0:
				dp[step][i+1] += dp[step-1][i] % MOD

			case i == 25 && dp[step-1][25] > 0:
				dp[step][0] = (dp[step][0] + dp[step-1][25]) % MOD
				dp[step][1] = (dp[step][1] + dp[step-1][25]) % MOD
			}
		}
	}

	for _, c := range dp[t] {
		cnt += c % MOD
	}

	return cnt % MOD
}
