/*
	https://leetcode.com/problems/optimal-partition-of-string/

	Given a string s, partition the string into one or more substrings such that the
	characters in each substring are unique. That is, no letter appears in a single
	substring more than once.

	Return the minimum number of substrings in such a partition.

	Note that each character should belong to exactly one substring in a partition.
*/

package solutions

func partitionString(s string) int {
	ans := 1
	mask := 0
	for _, ch := range s {
		shift := ch - 'a'
		if (mask & (1 << shift)) > 0 {
			ans++
			mask = 0
		}
		mask |= 1 << shift
	}
	return ans
}
