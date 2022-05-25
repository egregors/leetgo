/*
	https://leetcode.com/problems/russian-doll-envelopes/

	You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents
	the width and the height of an envelope.

	One envelope can fit into another if and only if both the width and height of one envelope
	are greater than the other envelope's width and height.

	Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).

	Note: You cannot rotate an envelope.
*/

package solutions

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var dp []int
	for _, envelope := range envelopes {
		i := sort.SearchInts(dp, envelope[1])
		if i == len(dp) {
			dp = append(dp, envelope[1])
		} else {
			dp[i] = envelope[1]
		}
	}

	return len(dp)
}
