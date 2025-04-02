/*
	https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/description/

		You are given a 0-indexed integer array nums.

	Return the maximum value over all triplets of indices (i, j, k) such

		that i < j < k. If all such triplets have a negative value, return 0.

	The value of a triplet of indices (i, j, k) is equal to
	(nums[i] - nums[j]) * nums[k].
*/

package solutions

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	var res, maxI, maxDiff int64 = 0, 0, 0
	for k := 0; k < n; k++ {
		res = max(res, maxDiff*int64(nums[k]))
		maxDiff = max(maxDiff, maxI-int64(nums[k]))
		maxI = max(maxI, int64(nums[k]))
	}

	return res
}
