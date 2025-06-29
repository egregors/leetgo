/*
	https://leetcode.com/problems/maximum-ascending-subarray-sum/

	Given an array of positive integers nums, return the maximum possible sum of an
	in nums.

	A subarray is defined as a contiguous sequence of numbers in an array.
*/

package solutions

import "math"

func maxAscendingSum(nums []int) int {
	prev := nums[0]
	acc := []int{prev}
	res := math.MinInt

	sum := func(xs ...int) int {
		s := 0
		for _, x := range xs {
			s += x
		}
		return s
	}

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n <= prev {
			res = max(res, sum(acc...))
			acc = []int{n}
		} else {
			acc = append(acc, n)
		}
		prev = n
	}

	return max(res, sum(acc...))
}
