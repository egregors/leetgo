/*
	https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/description/

		You are given a 0-indexed integer array nums.

	Return the maximum value over all triplets of indices (i, j, k) such that i < j
		< k.
	If all such triplets have a negative value, return 0.

	The value of a triplet of indices (i, j, k) is equal to (nums[i] - nums[j]) *
		nums[k].
*/

package solutions

func maximumTripletValue2874(nums []int) int64 {
	n := len(nums)
	res, maxL, maxD := 0, 0, 0
	for i := 0; i < n; i++ {
		res = max(res, maxD*nums[i])
		maxD = max(maxD, maxL-nums[i])
		maxL = max(maxL, nums[i])
	}

	return int64(res)
}
