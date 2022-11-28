/*
	https://leetcode.com/problems/find-players-with-zero-or-one-losses/

		You are given an integer array matches where matches[i] = [winneri, loseri]
	indicates that the player winneri defeated player loseri in a match.

	Return a list answer of size 2 where:

		answer[0] is a list of all players that have not lost any matches.
		answer[1] is a list of all players that have lost exactly one match.

	The values in the two lists should be returned in increasing order.

	Note:

		You should only consider the players that have played at least one match.
		The testcases will be generated such that no two matches will have the same
	outcome.
*/

package solutions

import "sort"

func findWinners(matches [][]int) [][]int {
	m := make(map[int]int)
	for _, match := range matches {
		if _, ok := m[match[0]]; !ok {
			m[match[0]] = 0
		}
		m[match[1]]++
	}

	var f, s []int
	for k, v := range m {
		if v == 0 {
			f = append(f, k)
		}
		if v == 1 {
			s = append(s, k)
		}
	}

	sort.Ints(f)
	sort.Ints(s)

	return [][]int{f, s}
}
