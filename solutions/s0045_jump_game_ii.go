/*
	https://leetcode.com/problems/jump-game-ii/

	You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

	Each element nums[i] represents the maximum length of a forward jump from index i. In other words,
	if you are at nums[i], you can jump to any nums[i + j] where:

		0 <= j <= nums[i] and
		i + j < n

	Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach
	nums[n - 1].
*/

package solutions

func jump(nums []int) int {
	ans, end, far := 0, 0, 0
	n := len(nums)

	for i := 0; i < n-1; i++ {
		far = Maximum(far, i+nums[i])
		if i == end {
			ans++
			end = far
		}
	}
	return ans
}
