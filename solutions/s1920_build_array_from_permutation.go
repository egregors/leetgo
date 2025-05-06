/*
	https://leetcode.com/problems/build-array-from-permutation/

	Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]]
	for each 0 <= i < nums.length and return it.

	A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).
*/

package solutions

func buildArray0005(nums []int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}
