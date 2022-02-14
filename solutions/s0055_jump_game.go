/*
	https://leetcode.com/problems/jump-game/

	You are given an integer array nums. You are initially positioned at the array's first
	index, and each element in the array represents your maximum jump length at that position.

	Return true if you can reach the last index, or false otherwise.

*/

package solutions

func canJump(nums []int) bool {
	UNKNOWN := 0
	GOOD := 1

	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = UNKNOWN
	}

	memo[len(memo)-1] = GOOD

	for i := len(nums) - 2; i >= 0; i-- {
		furthestJump := Minimum(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthestJump; j++ {
			if memo[j] == GOOD {
				memo[i] = GOOD
				break
			}
		}
	}
	return memo[0] == GOOD
}
