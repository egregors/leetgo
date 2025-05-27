/*
	https://leetcode.com/problems/orderly-queue/

	You are given a string s and an integer k. You can choose one of the first k
		letters
	of s and append it at the end of the string..

	Return the lexicographically smallest string you could have after applying the
		mentioned
	step any number of moves.
*/

package solutions

import (
	"sort"
	"strings"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		answer := s
		double := s + s
		for i := range s {
			temp := double[i : i+len(s)]
			if temp < answer {
				answer = temp
			}
		}
		return answer
	}
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}
