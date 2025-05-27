/*
	https://leetcode.com/problems/group-anagrams/

	Given an array of strings strs, group the anagrams together. You can return the
		answer in any order.

	An Anagram is a word or phrase formed by rearranging the letters of a different
		word or phrase,
	typically using all the original letters exactly once.
*/

package solutions

import "fmt"

func makeKey(s string) (res string) {
	m := make(map[int]int)
	for _, ch := range s {
		m[int(ch-'a')]++
	}

	for i := 0; i < 26; i++ {
		res += fmt.Sprintf("%d.", m[i])
	}

	return res[:len(res)-1]
}

func groupAnagrams(strs []string) [][]string {
	ws := make(map[string][]string)
	for _, s := range strs {
		k := makeKey(s)
		ws[k] = append(ws[k], s)
	}

	var res [][]string //nolint:prealloc //meh
	for _, v := range ws {
		res = append(res, v)
	}

	return res
}
