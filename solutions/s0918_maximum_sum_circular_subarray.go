/*
	https://leetcode.com/problems/maximum-sum-circular-subarray/

	Given a circular integer array nums of length n, return the maximum possible
		sum of a
	non-empty subarray of nums.

	A circular array means the end of the array connects to the beginning of the
		array.
	Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous
		element
	of nums[i] is nums[(i - 1 + n) % n].

	A subarray may only include each element of the fixed buffer nums at most once.
	Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not
		exist
	i <= k1, k2 <= j with k1 % n == k2 % n.
*/

package solutions

func maxSubarraySumCircular(nums []int) int {
	curMax, curMin, sum := 0, 0, 0
	maxSum, minSum := nums[0], nums[0]
	for _, n := range nums {
		curMax = Maximum(curMax, 0) + n
		maxSum = Maximum(maxSum, curMax)
		curMin = Minimum(curMin, 0) + n
		minSum = Minimum(minSum, curMin)
		sum += n
	}
	if sum == minSum {
		return maxSum
	}
	return Maximum(maxSum, sum-minSum)
}
