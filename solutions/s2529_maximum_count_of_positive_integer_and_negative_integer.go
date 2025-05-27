/*
	https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/

	Given an array nums sorted in non-decreasing order, return the maximum between
		the number
	of positive integers and the number of negative integers.

    In other words, if the number of positive integers in nums is pos and the
    	number of negative
	integers is neg, then return the maximum of pos and neg.

	Note that 0 is neither positive nor negative.
*/

package solutions

func maximumCount(nums []int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo <= hi {
		mid = (lo + hi) >> 1
		if nums[mid] > 0 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	i := mid
	for i >= 0 && nums[i] >= 0 {
		i--
	}
	neg := i + 1
	i = mid
	for i < len(nums) && nums[i] <= 0 {
		i++
	}
	pos := len(nums) - i

	if neg > pos {
		return neg
	}

	return pos
}
