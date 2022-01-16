/*
	https://leetcode.com/problems/partition-labels/

	You are given a string s. We want to partition the string into as many parts as possible
	so that each letter appears in at most one part.

	Note that the partition is done so that after concatenating all the parts in order, the
	resultant string should be s.

	Return a list of integers representing the size of these parts.
*/

package solutions

func partitionLabels(s string) []int {
	last := make([]int, 26)

	for i, ch := range s {
		last[ch-'a'] = i
	}

	var j, anchor int
	var res []int
	for i, ch := range s {
		j = max(j, last[ch-'a'])
		if i == j {
			res = append(res, i-anchor+1)
			anchor = i + 1
		}
	}

	return res
}
