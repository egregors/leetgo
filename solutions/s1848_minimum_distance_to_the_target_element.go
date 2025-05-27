/*
	https://leetcode.com/problems/minimum-distance-to-the-target-element

	Given an integer array nums (0-indexed) and two integers target and start, find
		an index i such that
	nums[i] == target and abs(i - start) is minimized. Note that abs(x) is the
		absolute value of x.

	Return abs(i - start).

	It is guaranteed that target exists in nums.
*/

package solutions

import "math"

func getMinDistance(nums []int, target int, start int) int {
	val := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			val = min(val, abs(i-start))
		}
	}

	return val
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}
