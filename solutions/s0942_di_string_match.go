/*
	https://leetcode.com/problems/di-string-match/

	A permutation perm of n + 1 integers of all the integers in the range [0, n]
	can be represented as a string s of length n where:

		s[i] == 'I' if perm[i] < perm[i + 1], and
		s[i] == 'D' if perm[i] > perm[i + 1].

	Given a string s, reconstruct the permutation perm and return it.
	If there are multiple valid permutations perm, return any of them.
*/

package solutions

func diStringMatch(s string) []int {
	n := len(s)
	lo, hi := 0, n
	ans := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ans[i] = lo
			lo++
		} else {
			ans[i] = hi
			hi--
		}
	}
	ans[n] = lo
	return ans
}
