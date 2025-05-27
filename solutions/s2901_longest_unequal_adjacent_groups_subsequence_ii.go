/*
	https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/

	You are given a string array words, and an array groups, both arrays having
		length n.

	The hamming distance between two strings of equal length is the number of
		positions at which the corresponding
	characters are different.

	You need to select the longest

	from an array of indices [0, 1, ..., n - 1], such that for the subsequence
		denoted as [i0, i1, ..., ik-1] having
	length k, the following holds:

		For adjacent indices in the subsequence, their corresponding groups are
			unequal, i.e.,
	groups[ij] != groups[ij+1], for each j where 0 < j + 1 < k.
		words[ij] and words[ij+1] are equal in length, and the hamming distance
			between them is 1, where
	0 < j + 1 < k, for all indices in the subsequence.

	Return a string array containing the words corresponding to the indices (in
		order) in the selected subsequence.
	If there are multiple answers, return any of them.

	Note: strings in words may be unequal in length.
*/

package solutions

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxIdx := 0

	pred := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		diff := 0
		for i := range len(a) {
			if a[i] != b[i] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
		return diff == 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if pred(words[i], words[j]) && dp[j]+1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}

	ans := []string{}
	for i := maxIdx; i >= 0; i = prev[i] {
		ans = append(ans, words[i])
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return ans
}
