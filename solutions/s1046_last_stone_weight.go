/*
	https://leetcode.com/problems/last-stone-weight/

	You are given an array of integers stones where stones[i] is the weight of the
		ith stone.

	We are playing a game with the stones. On each turn, we choose the heaviest two
		stones and
	smash them together. Suppose the heaviest two stones have weights x and y with
		x <= y.
	The result of this smash is:

		If x == y, both stones are destroyed, and
		If x != y, the stone of weight x is destroyed, and the stone of weight y has
			new
	weight y - x.

	At the end of the game, there is at most one stone left.

	Return the smallest possible weight of the left stone. If there are no stones
		left, return 0.
*/

package solutions

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
		a, b := stones[0], stones[1]
		if a == b {
			stones = stones[2:]
		} else {
			left := a - b
			stones = stones[1:]
			stones[0] = left
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
