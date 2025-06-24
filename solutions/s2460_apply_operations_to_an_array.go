/*
	https://leetcode.com/problems/apply-operations-to-an-array

	You are given a 0-indexed array nums of size n consisting of non-negative
		integers.

	You need to apply n - 1 operations to this array where, in the ith operation
		(0-indexed), you will apply the following on the ith element of nums:

		If nums[i] == nums[i + 1], then multiply nums[i] by 2 and set nums[i + 1] to
			0. Otherwise, you skip this operation.

	After performing all the operations, shift all the 0's to the end of the array.

		For example, the array [1,0,2,0,0,1] after shifting all its 0's to the end, is
			[1,2,1,0,0,0].

	Return the resulting array.

	Note that the operations are applied sequentially, not all at once.
*/

package solutions

func applyOperations(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums[i] = 0
			nums[i-1] *= 2
		}
	}

	res := make([]int, 0, len(nums))
	for _, n := range nums {
		if n != 0 {
			res = append(res, n)
		}
	}

	return res[0:len(nums):len(nums)]
}
