/*
	https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/

	You are given an array of strings arr. A string s is formed by the concatenation of a subsequence
	of arr that has unique characters.

	Return the maximum possible length of s.

	A subsequence is an array that can be derived from another array by deleting some or no elements
	without changing the order of the remaining elements.
*/

package solutions

func maxLength(arr []string) int {
	var dfs func(arr []string, pos int, res string) int
	dfs = func(arr []string, pos int, res string) int {
		if len(res) != len(NewSet([]rune(res)).ToSlice()) {
			return 0
		}
		best := len(res)
		for i := pos; i < len(arr); i++ {
			best = Maximum(best, dfs(arr, i+1, res+arr[i]))
		}
		return best
	}
	return dfs(arr, 0, "")
}
