/*
	https://leetcode.com/problems/house-robber/

	You are a professional robber planning to rob houses along a street.
	Each house has a certain amount of money stashed, the only constraint stopping
		you
	from robbing each of them is that adjacent houses have security systems
		connected and
	it will automatically contact the police if two adjacent houses were broken
		into on the
	same night.

	Given an integer array nums representing the amount of money of each house,
	return the maximum amount of money you can rob tonight without alerting the
		police.
*/

package solutions

func dp198(i int, nums []int, mem map[int]int) int {
	switch i {
	case 0:
		return nums[0]
	case 1:
		return Maximum(nums[0], nums[1])
	default:
		if _, ok := mem[i]; !ok {
			mem[i] = Maximum(dp198(i-1, nums, mem), dp198(i-2, nums, mem)+nums[i])
		}
		return mem[i]
	}
}

func rob(nums []int) int {
	return dp198(len(nums)-1, nums, make(map[int]int))
}
