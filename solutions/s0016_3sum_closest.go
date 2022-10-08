/*
	https://leetcode.com/problems/3sum-closest/

	Given an integer array nums of length n and an integer target, find three integers
	in nums such that the sum is closest to target.

	Return the sum of the three integers.

	You may assume that each input would have exactly one solution.
*/

package solutions

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt
	sz := len(nums)
	sort.Ints(nums)
	for i := 0; i < sz && diff != 0; i++ {
		lo, hi := i+1, sz-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if Abs(target-sum) < Abs(diff) {
				diff = target - sum
			}
			if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return target - diff
}
