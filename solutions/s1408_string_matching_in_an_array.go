/*
	https://leetcode.com/problems/string-matching-in-an-array

	Given an array of string words, return all strings in words that are a of
		another word. You can return the answer in any order.
*/

package solutions

import (
	"maps"
	"slices"
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := make(map[string]bool)
	for i := len(words) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if strings.Contains(words[i], words[j]) {
				res[words[j]] = true
			}
		}
	}

	return slices.Collect(maps.Keys(res))
}
