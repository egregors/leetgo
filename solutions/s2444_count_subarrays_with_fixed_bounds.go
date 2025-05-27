/*
	https://leetcode.com/problems/count-subarrays-with-fixed-bounds/

	You are given an integer array nums and two integers minK and maxK.

	A fixed-bound subarray of nums is a subarray that satisfies the following
		conditions:

		The minimum value in the subarray is equal to minK.
		The maximum value in the subarray is equal to maxK.

	Return the number of fixed-bound subarrays.

	A subarray is a contiguous part of an array.
*/

package solutions

func countSubarrays(nums []int, minK, maxK int) int64 {
	ans := 0
	minP, maxP, leftB := -1, -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			leftB = i
		}

		if nums[i] == minK {
			minP = i
		}
		if nums[i] == maxK {
			maxP = i
		}

		ans += Maximum(0, Minimum(maxP, minP)-leftB)
	}
	return int64(ans)
}
