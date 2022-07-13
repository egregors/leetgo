/*
	https://leetcode.com/problems/matchsticks-to-square/

	You are given an integer array matchsticks where matchsticks[i] is the length of
	the ith matchstick. You want to use all the matchsticks to make one square.
	You should not break any stick, but you can link them up, and each matchstick
	must be used exactly one time.

	Return true if you can make this square and false otherwise.
*/

package solutions

import "sort"

func makesquare(matchsticks []int) bool {
	sum := Sum(matchsticks...)
	if sum%4 != 0 {
		return false
	}
	sideLen, sides := sum/4, make([]int, 4)
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	var bt func(i int) bool
	bt = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}
		for j := 0; j < 4; j++ {
			if sides[j]+matchsticks[i] <= sideLen {
				sides[j] += matchsticks[i]
				if bt(i + 1) {
					return true
				}
				sides[j] -= matchsticks[i]
			}
		}
		return false
	}
	return bt(0)
}
