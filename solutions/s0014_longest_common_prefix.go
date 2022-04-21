/*
	https://leetcode.com/problems/longest-common-prefix/

	Write a function to find the longest common prefix string amongst an array of strings.

	If there is no common prefix, return an empty string "".
*/

package solutions

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	if len(strs) == 1 {
		return strs[0]
	}

	l := 0
	for c := 0; c < len(strs[0]); c++ {
		for r := 1; r < len(strs); r++ {
			if strs[r][c] != strs[0][c] {
				return strs[0][:l]
			}
		}
		l++
	}

	return strs[0][:l]
}
