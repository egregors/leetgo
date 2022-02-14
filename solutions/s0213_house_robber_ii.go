/*
	https://leetcode.com/problems/house-robber-ii/

	You are a professional robber planning to rob houses along a street.
	Each house has a certain amount of money stashed. All houses at this place are arranged
	in a circle. That means the first house is the neighbor of the last one. Meanwhile,
	adjacent houses have a security system connected, and it will automatically contact the
	police if two adjacent houses were broken into on the same night.

	Given an integer array nums representing the amount of money of each house, return
	the maximum amount of money you can rob tonight without alerting the police.
*/

package solutions

// rob213 must call rob to pass LeetCode tests
func rob213(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	maxRobbedAmountA := make([]int, n+1)
	maxRobbedAmountB := make([]int, n+1)

	numsA := nums[:len(nums)-1]
	numsB := nums[1:]
	n--

	// base
	maxRobbedAmountA[n] = 0
	maxRobbedAmountA[n-1] = numsA[n-1]
	maxRobbedAmountB[n] = 0
	maxRobbedAmountB[n-1] = numsB[n-1]

	// DP table
	for i := n - 2; i >= 0; i-- {
		maxRobbedAmountA[i] = Maximum(maxRobbedAmountA[i+1], maxRobbedAmountA[i+2]+nums[i])
		maxRobbedAmountB[i] = Maximum(maxRobbedAmountB[i+1], maxRobbedAmountB[i+2]+numsB[i])
	}

	if maxRobbedAmountA[0] > maxRobbedAmountB[0] {
		return maxRobbedAmountA[0]
	}
	return maxRobbedAmountB[0]
}
