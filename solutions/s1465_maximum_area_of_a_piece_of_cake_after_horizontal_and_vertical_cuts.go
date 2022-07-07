/*
	https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/

		You are given a rectangular cake of size h x w and two arrays of integers horizontalCuts
		and verticalCuts where:

		horizontalCuts[i] is the distance from the top of the rectangular cake to the ith horizontal
		cut and similarly, and
		verticalCuts[j] is the distance from the left of the rectangular cake to the jth vertical
		cut.

	Return the maximum area of a piece of cake after you cut at each horizontal and vertical
	position provided in the arrays horizontalCuts and verticalCuts. Since the answer can be a
	large number, return this modulo 109 + 7.
*/

package solutions

import "sort"

// maxArea1465 should call maxArea to pass LeetCode tests
func maxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	n := len(horizontalCuts)
	m := len(verticalCuts)

	maxH := Maximum(horizontalCuts[0], h-horizontalCuts[n-1])
	for i := 1; i < n; i++ {
		maxH = Maximum(maxH, horizontalCuts[i]-horizontalCuts[i-1])
	}

	maxW := Maximum(verticalCuts[0], w-verticalCuts[m-1])
	for i := 1; i < m; i++ {
		maxW = Maximum(maxW, verticalCuts[i]-verticalCuts[i-1])
	}

	return (maxW * maxH) % 1000000007
}
